package http

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"icmongolang/config"
	"icmongolang/internal/auth"
	"icmongolang/internal/middleware"
	"icmongolang/internal/users"
	"icmongolang/internal/users/presenter"
	"icmongolang/pkg/httpErrors"
	"icmongolang/pkg/jwt"
	"icmongolang/pkg/logger"
	"icmongolang/pkg/responses"
	"icmongolang/pkg/utils"

	"github.com/go-chi/render"
)

type authHandler struct {
	usersUC             users.UserUseCaseI
	cfg                 *config.Config
	logger              logger.Logger
	rateLimitMiddleware *middleware.RateLimitMiddleware
}

// CreateAuthHandler creates auth handler instance
func CreateAuthHandler(uc users.UserUseCaseI, cfg *config.Config, logger logger.Logger) auth.Handlers {
	// สร้าง rate limit middleware (กำหนดค่าเองได้)
	rateLimitConfig := middleware.RateLimitConfig{
		// RequestsPerSecond: จำนวนคำขอสูงสุดที่อนุญาตต่อวินาที
		// Maximum number of requests allowed per second
		RequestsPerSecond: 50, // 50 requests per second | 50 คำขอต่อวินาที

		// Burst: จำนวนคำขอสูงสุดที่อนุญาตให้ส่งพร้อมกัน (กระชุ)
		// Maximum number of requests allowed to burst at once
		Burst: 100, // Burst up to 100 | อนุญาตให้ส่งสูงสุด 100 คำขอในครั้งเดียว

		// CleanupInterval: ระยะเวลาในการทำความสะอาดข้อมูล IP ที่ไม่ใช้งาน
		// Interval for cleaning up inactive IP data
		CleanupInterval: 15 * time.Minute, // Clean up every 15 minutes | ทำความสะอาดทุก 15 นาที
	}
	rateLimitMiddleware := middleware.NewRateLimitMiddleware(rateLimitConfig)

	return &authHandler{
		cfg:                 cfg,
		usersUC:             uc,
		logger:              logger,
		rateLimitMiddleware: rateLimitMiddleware, // เพิ่ม
	}
}

// SignIn godoc
// @Summary Sign In
// @Description Sign in, get access token for future requests.
// @Tags auth
// @Accept multipart/form-data
// @Produce json
// @Param username formData string true "email"
// @Param password formData string true "password"
// @Success 200 {object} presenter.Token
// @Failure 400	{object} responses.ErrorResponse
// @Failure 422	{object} responses.ErrorResponse
// @Router /auth/login [post]
func (h *authHandler) SignIn() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// ใช้ rate limit middleware ตรวจสอบก่อน
		// สร้าง handler ชั่วคราวเพื่อตรวจสอบ rate limit
		checkHandler := h.rateLimitMiddleware.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// ถ้าผ่าน rate limit ให้ทำงานปกติ
		}))

		// สร้าง response writer จับ error
		recorder := &responseRecorder{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		// ตรวจสอบ rate limit
		checkHandler.ServeHTTP(recorder, r)

		// ถ้า status เป็น 429 แสดงว่าโดน rate limit
		if recorder.statusCode == http.StatusTooManyRequests {
			return
		}

		// Log: เริ่มต้นการทำงาน
		// Log: Start processing
		h.logger.Info("========== SignIn request received ==========")

		// ประกาศตัวแปรสำหรับรับข้อมูลผู้ใช้
		// Declare variable for user data
		user := new(presenter.UserSignIn)

		// ตรวจสอบ Content-Type จาก Header
		// Check Content-Type from Header
		contentType := r.Header.Get("Content-Type")
		h.logger.Infof("Content-Type: %s", contentType) // Log Content-Type

		var email, password string

		// กรณีรับข้อมูลแบบ JSON
		// Case: receive data as JSON
		if strings.Contains(contentType, "application/json") {
			h.logger.Info("Processing JSON request") // Log: กำลังประมวลผล JSON

			// อ่าน Body ทั้งหมด
			// Read entire body
			body, err := io.ReadAll(r.Body)
			if err != nil {
				h.logger.Errorf("Failed to read body: %v", err) // Log error
				render.Render(w, r, responses.CreateErrorResponse(err))
				return
			}
			defer r.Body.Close()

			h.logger.Debugf("Raw JSON body: %s", string(body)) // Log: แสดง JSON ที่รับมา

			// แปลง JSON เป็น map
			// Parse JSON to map
			var jsonBody map[string]interface{}
			if err := json.Unmarshal(body, &jsonBody); err != nil {
				h.logger.Errorf("Failed to parse JSON: %v", err) // Log error
				render.Render(w, r, responses.CreateErrorResponse(httpErrors.ErrValidation(err)))
				return
			}

			// ดึงค่า username และ password จาก JSON
			// Extract username and password from JSON
			email, _ = jsonBody["username"].(string)
			password, _ = jsonBody["password"].(string)

			h.logger.Debugf("Extracted from JSON - Email: %s, Password length: %d", email, len(password)) // Log: ค่าที่ดึงได้
		} else {
			// กรณีรับข้อมูลแบบ x-www-form-urlencoded (แบบเดิม)
			// Case: receive data as x-www-form-urlencoded (original way)
			h.logger.Info("Processing form-urlencoded request") // Log: กำลังประมวลผล Form

			if err := r.ParseForm(); err != nil {
				h.logger.Errorf("Failed to parse form: %v", err) // Log error
				render.Render(w, r, responses.CreateErrorResponse(err))
				return
			}
			// ดึงค่าจากฟอร์ม
			// Get values from form
			email = r.FormValue("username")
			password = r.FormValue("password")

			h.logger.Debugf("Extracted from Form - Email: %s, Password length: %d", email, len(password)) // Log: ค่าที่ดึงได้
		}

		// ตรวจสอบว่ามีค่าหรือไม่
		// Check if values exist
		if email == "" {
			h.logger.Warn("Email is empty") // Log warning
		}
		if password == "" {
			h.logger.Warn("Password is empty") // Log warning
		}

		// กำหนดค่าให้กับ struct user
		// Assign values to user struct
		user.Email = email
		user.Password = password

		h.logger.Infof("Attempting login with email: %s", email) // Log: กำลังพยายาม login

		// ตรวจสอบความถูกต้องของข้อมูล (required, format ฯลฯ)
		// Validate data (required, format, etc.)
		err := utils.ValidateStruct(r.Context(), user)
		if err != nil {
			h.logger.Errorf("Validation failed: %v", err) // Log error validation
			render.Render(w, r, responses.CreateErrorResponse(httpErrors.ErrValidation(err)))
			return
		}

		h.logger.Info("Validation passed, calling SignIn usecase") // Log: ผ่าน validation

		// เรียกใช้ UseCase เพื่อทำการ Sign In
		// Call UseCase to perform Sign In
		accessToken, refreshToken, err := h.usersUC.SignIn(
			r.Context(),
			user.Email,
			user.Password,
		)
		if err != nil {
			h.logger.Errorf("SignIn failed: %v", err) // Log error จาก usecase
			render.Render(w, r, responses.CreateErrorResponse(err))
			return
		}

		// คำนวณเวลาหมดอายุของ token (หน่วยเป็นวินาที)
		// Calculate token expiration time (in seconds)
		// สมมติว่า config เก็บเป็นนาที
		// Assume config stores value in minutes
		expiresIn := h.cfg.Jwt.AccessTokenExpireDuration * 60

		h.logger.Info("SignIn successful, returning tokens")                                                      // Log: สำเร็จ
		h.logger.Debugf("Access token length: %d, Refresh token length: %d", len(accessToken), len(refreshToken)) // Log: ความยาว token
		h.logger.Info("========== SignIn request completed ==========")

		// ส่ง response กลับไปพร้อม access token และ refresh token
		// Send response back with access token and refresh token
		render.Respond(w, r, map[string]interface{}{
			"access_token":  accessToken,  // โทเคนสำหรับเข้าถึง API / Token for API access
			"refresh_token": refreshToken, // โทเคนสำหรับขอ access token ใหม่ / Token for refreshing access token
			"token_type":    "bearer",     // ประเภทของโทเคน / Type of token
			"expires_in":    expiresIn,    // อายุของโทเคน (วินาที) / Token lifetime (seconds)
		})
	}
}

// responseRecorder สำหรับจับ status code
type responseRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rr *responseRecorder) WriteHeader(code int) {
	rr.statusCode = code
	rr.ResponseWriter.WriteHeader(code)
}

func (rr *responseRecorder) Write(b []byte) (int, error) {
	return rr.ResponseWriter.Write(b)
}

// RefreshToken godoc
// @Summary Refresh token
// @Description Get new access token from refresh token.
// @Tags auth
// @Accept json
// @Produce json
// @Param Authorization header string true "Authentication header"
// @Success 200 {object} presenter.Token
// @Failure 400	{object} responses.ErrorResponse
// @Failure 422	{object} responses.ErrorResponse
// @Router /auth/refresh [get]
func (h *authHandler) RefreshToken() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		refreshToken := middleware.TokenFromHeader(r)

		accessToken, refreshToken, err := h.usersUC.Refresh(ctx, refreshToken)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err)) //nolint:errcheck
			return
		}

		render.Respond(w, r, presenter.Token{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			TokenType:    "bearer",
		})
	}
}

// GetPublicKey godoc
// @Summary Get public key
// @Description Get rsa public key to decode token.
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} presenter.PublicKey
// @Failure 400	{object} responses.ErrorResponse
// @Failure 422	{object} responses.ErrorResponse
// @Router /auth/publickey [get]
func (h *authHandler) GetPublicKey() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		publicKeyAccessToken, err := jwt.DecodeBase64(h.cfg.Jwt.AccessTokenPublicKey)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err)) //nolint:errcheck
			return
		}

		publicKeyRefreshToken, err := jwt.DecodeBase64(h.cfg.Jwt.RefreshTokenPublicKey)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err)) //nolint:errcheck
			return
		}

		render.Respond(w, r, presenter.PublicKey{
			PublicKeyAccessToken:  string(publicKeyAccessToken[:]),
			PublicKeyRefreshToken: string(publicKeyRefreshToken[:]),
		})
	}
}

// Logout godoc
// @Summary Logout
// @Description Logout, remove current refresh token in db.
// @Tags auth
// @Accept json
// @Produce json
// @Param Authorization header string true "Authentication header"
// @Success 200
// @Failure 400	{object} responses.ErrorResponse
// @Failure 422	{object} responses.ErrorResponse
// @Router /auth/logout [get]
func (h *authHandler) Logout() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		refreshToken := middleware.TokenFromHeader(r)

		err := h.usersUC.Logout(ctx, refreshToken)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err)) //nolint:errcheck
			return
		}
	}
}

// LogoutAllToken godoc
// @Summary Logout all session
// @Description Logout all session of this user.
// @Tags auth
// @Accept json
// @Produce json
// @Param Authorization header string true "Authentication header"
// @Success 200
// @Failure 400	{object} responses.ErrorResponse
// @Failure 422	{object} responses.ErrorResponse
// @Router /auth/logoutall [get]
func (h *authHandler) LogoutAllToken() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		refreshToken := middleware.TokenFromHeader(r)

		id, err := h.usersUC.ParseIdFromRefreshToken(ctx, refreshToken)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err)) //nolint:errcheck
			return
		}

		err = h.usersUC.LogoutAll(ctx, id)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err)) //nolint:errcheck
			return
		}
	}
}

// VerifyEmail godoc
// @Summary Verify user
// @Description Verify user using code from email.
// @Tags auth
// @Accept json
// @Produce json
// @Param code query string true "offset" Format(code)
// @Success 200 {object} string
// @Failure 400	{object} responses.ErrorResponse
// @Router /auth/verifyemail [get]
func (h *authHandler) VerifyEmail() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		q := r.URL.Query()
		verificationCode := q.Get("code")

		err := h.usersUC.Verify(ctx, verificationCode)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err)) //nolint:errcheck
			return
		}

		render.Respond(w, r, responses.CreateSuccessResponse("Email verified successfully"))
	}
}

// ForgotPassword godoc
// @Summary Forgot password
// @Description Forgot password, code will send to email.
// @Tags auth
// @Accept json
// @Produce json
// @Param forgotPassword body presenter.ForgotPassword true "Forgot Password"
// @Success 200 {object} string
// @Failure 400	{object} responses.ErrorResponse
// @Failure 422	{object} responses.ErrorResponse
// @Router /auth/forgotpassword [post]
func (h *authHandler) ForgotPassword() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		forgotPassword := new(presenter.ForgotPassword)

		err := json.NewDecoder(r.Body).Decode(&forgotPassword)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err)) //nolint:errcheck
			return
		}

		err = utils.ValidateStruct(r.Context(), forgotPassword)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(httpErrors.ErrValidation(err))) //nolint:errcheck
			return
		}

		err = h.usersUC.ForgotPassword(ctx, forgotPassword.Email)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err)) //nolint:errcheck
			return
		}

		render.Respond(w, r,
			responses.CreateSuccessResponse("You will receive a reset email if user with that email exist"))
	}
}

// ResetPassword godoc
// @Summary Reset Password
// @Description Reset Password, using code from email.
// @Tags auth
// @Accept json
// @Produce json
// @Param code query string true "code" Format(code)
// @Param resetPassword body presenter.ResetPassword true "Reset Password"
// @Success 200 {object} string
// @Failure 400	{object} responses.ErrorResponse
// @Failure 422	{object} responses.ErrorResponse
// @Router /auth/resetpassword [patch]
func (h *authHandler) ResetPassword() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		q := r.URL.Query()
		resetToken := q.Get("code")

		resetPassword := new(presenter.ResetPassword)

		err := json.NewDecoder(r.Body).Decode(&resetPassword)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err)) //nolint:errcheck
			return
		}

		err = utils.ValidateStruct(r.Context(), resetPassword)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(httpErrors.ErrValidation(err))) //nolint:errcheck
			return
		}

		err = h.usersUC.ResetPassword(
			ctx,
			resetToken,
			resetPassword.NewPassword,
			resetPassword.ConfirmPassword,
		)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err)) //nolint:errcheck
			return
		}

		render.Respond(w, r,
			responses.CreateSuccessResponse("Password data updated successfully, please re-login"))
	}
}
