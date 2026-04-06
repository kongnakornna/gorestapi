package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"gorestapi/config"
	"gorestapi/internal/middleware"
	"gorestapi/internal/models"
	"gorestapi/internal/users"
	"gorestapi/internal/users/presenter"
	"gorestapi/pkg/httpErrors"
	"gorestapi/pkg/logger"
	"gorestapi/pkg/responses"
	"gorestapi/pkg/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type userHandler struct {
	cfg     *config.Config
	usersUC users.UserUseCaseI
	logger  logger.Logger
}



// CreateUserHandler สร้าง instance ของ user handler
// CreateUserHandler creates a new user handler instance
func CreateUserHandler(uc users.UserUseCaseI, cfg *config.Config, logger logger.Logger) users.Handlers {
	return &userHandler{cfg: cfg, usersUC: uc, logger: logger}
}

// ==================== EXISTING HANDLERS (Protected / Admin) ====================

// Create สร้างผู้ใช้ใหม่ (เฉพาะ superuser)
// Create creates a new user (superuser only)
// @Summary Create User
// @Description Create new user.
// @Tags users
// @Accept json
// @Produce json
// @Param user body presenter.UserCreate true "Add user"
// @Success 200 {object} responses.SuccessResponse[presenter.UserResponse]
// @Failure 400 {object} responses.ErrorResponse
// @Security OAuth2Password
// @Router /user [post]
func (h *userHandler) Create() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user := new(presenter.UserCreate)

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err))
			return
		}

		err = utils.ValidateStruct(r.Context(), user)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(httpErrors.ErrValidation(err)))
			return
		}

		newUser, err := h.usersUC.CreateUser(
			r.Context(),
			mapModel(user),
			user.ConfirmPassword,
		)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err))
			return
		}

		userResponse := *mapModelResponse(newUser)
		render.Respond(w, r, responses.CreateSuccessResponse(userResponse))
	}
}

// Get ดึงข้อมูลผู้ใช้ตาม ID (ต้องเป็น superuser หรือเจ้าของตัวเอง)
// Get returns a user by ID (superuser or owner)
// @Summary Read user
// @Description Get user by ID.
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User Id"
// @Success 200 {object} responses.SuccessResponse[presenter.UserResponse]
// @Router /user/{id} [get]
func (h *userHandler) Get() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.Parse(chi.URLParam(r, "id"))
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(httpErrors.ErrValidation(err)))
			return
		}

		user, err := h.usersUC.Get(r.Context(), id)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err))
			return
		}

		render.Respond(w, r, responses.CreateSuccessResponse(mapModelResponse(user)))
	}
}

// GetMulti ดึงรายการผู้ใช้แบบ paginate (เฉพาะ superuser)
// GetMulti returns paginated users (superuser only)
// @Summary Read Users
// @Description Retrieve users.
// @Tags users
// @Accept json
// @Produce json
// @Param limit query int false "limit"
// @Param offset query int false "offset"
// @Success 200 {object} responses.SuccessResponse[[]presenter.UserResponse]
// @Router /user [get]
func (h *userHandler) GetMulti() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()

		limit, _ := strconv.Atoi(q.Get("limit"))
		offset, _ := strconv.Atoi(q.Get("offset"))

		users, err := h.usersUC.GetMulti(r.Context(), limit, offset)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err))
			return
		}

		render.Respond(w, r, responses.CreateSuccessResponse(mapModelsResponse(users)))
	}
}

// Delete ลบผู้ใช้ตาม ID (เฉพาะ superuser)
// Delete deletes a user by ID (superuser only)
// @Summary Delete user
// @Description Delete an user by ID.
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User Id"
// @Success 200 {object} responses.SuccessResponse[presenter.UserResponse]
// @Router /user/{id} [delete]
func (h *userHandler) Delete() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.Parse(chi.URLParam(r, "id"))
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(httpErrors.ErrValidation(err)))
			return
		}

		user, err := h.usersUC.Delete(r.Context(), id)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err))
			return
		}

		render.Respond(w, r, responses.CreateSuccessResponse(mapModelResponse(user)))
	}
}

// Update อัปเดตข้อมูลผู้ใช้ตาม ID (เฉพาะ superuser)
// Update updates a user by ID (superuser only)
// @Summary Update user
// @Description Update an user by ID.
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User Id"
// @Param user body presenter.UserUpdate true "Update user"
// @Success 200 {object} responses.SuccessResponse[presenter.UserResponse]
// @Router /user/{id} [put]
func (h *userHandler) Update() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.Parse(chi.URLParam(r, "id"))
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(httpErrors.ErrValidation(err)))
			return
		}

		user := new(presenter.UserUpdate)

		err = json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err))
			return
		}

		err = utils.ValidateStruct(r.Context(), user)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(httpErrors.ErrValidation(err)))
			return
		}

		values := make(map[string]interface{})
		if user.Name != "" {
			values["name"] = user.Name
		}

		updatedUser, err := h.usersUC.Update(r.Context(), id, values)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err))
			return
		}

		render.Respond(w, r, responses.CreateSuccessResponse(mapModelResponse(updatedUser)))
	}
}

// UpdatePassword อัปเดตรหัสผ่านผู้ใช้ตาม ID (เฉพาะ superuser)
// UpdatePassword updates user's password by ID (superuser only)
// @Summary Update password user
// @Description Update password user by ID.
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User Id"
// @Param user body presenter.UserUpdatePassword true "Update user"
// @Success 200 {object} responses.SuccessResponse[presenter.UserResponse]
// @Router /user/{id}/updatepass [patch]
func (h *userHandler) UpdatePassword() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.Parse(chi.URLParam(r, "id"))
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(httpErrors.ErrValidation(err)))
			return
		}

		user := new(presenter.UserUpdatePassword)

		err = json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err))
			return
		}

		err = utils.ValidateStruct(r.Context(), user)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(httpErrors.ErrValidation(err)))
			return
		}

		updatedUser, err := h.usersUC.UpdatePassword(
			r.Context(),
			id,
			user.OldPassword,
			user.NewPassword,
			user.ConfirmPassword,
		)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err))
			return
		}

		render.Respond(w, r, responses.CreateSuccessResponse(mapModelResponse(updatedUser)))
	}
}

// Me ดึงข้อมูลผู้ใช้ที่กำลัง login
// Me returns the currently logged-in user
// @Summary Read user me
// @Description Get user me.
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} responses.SuccessResponse[presenter.UserResponse]
// @Router /user/me [get]
func (h *userHandler) Me() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		user, err := middleware.GetUserFromCtx(ctx)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err))
			return
		}

		render.Respond(w, r, responses.CreateSuccessResponse(mapModelResponse(user)))
	}
}

// UpdateMe อัปเดตข้อมูลผู้ใช้ที่กำลัง login
// UpdateMe updates the currently logged-in user
// @Summary Update user me
// @Description Update user me.
// @Tags users
// @Accept json
// @Produce json
// @Param user body presenter.UserUpdate true "Update user"
// @Success 200 {object} responses.SuccessResponse[presenter.UserResponse]
// @Router /user/me [put]
func (h *userHandler) UpdateMe() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		user, err := middleware.GetUserFromCtx(ctx)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err))
			return
		}

		userUpdate := new(presenter.UserUpdate)

		err = json.NewDecoder(r.Body).Decode(&userUpdate)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err))
			return
		}

		err = utils.ValidateStruct(r.Context(), userUpdate)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(httpErrors.ErrValidation(err)))
			return
		}

		values := make(map[string]interface{})
		if userUpdate.Name != "" {
			values["name"] = userUpdate.Name
		}

		updatedUser, err := h.usersUC.Update(r.Context(), user.Id, values)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err))
			return
		}

		render.Respond(w, r, responses.CreateSuccessResponse(mapModelResponse(updatedUser)))
	}
}

// UpdatePasswordMe อัปเดตรหัสผ่านผู้ใช้ที่กำลัง login
// UpdatePasswordMe updates the currently logged-in user's password
// @Summary Update password user me
// @Description Update password user me.
// @Tags users
// @Accept json
// @Produce json
// @Param user body presenter.UserUpdatePassword true "Update user"
// @Success 200 {object} responses.SuccessResponse[presenter.UserResponse]
// @Router /user/me/updatepass [patch]
func (h *userHandler) UpdatePasswordMe() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		user, err := middleware.GetUserFromCtx(ctx)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err))
			return
		}

		userUpdate := new(presenter.UserUpdatePassword)

		err = json.NewDecoder(r.Body).Decode(&userUpdate)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err))
			return
		}

		err = utils.ValidateStruct(r.Context(), userUpdate)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(httpErrors.ErrValidation(err)))
			return
		}

		updatedUser, err := h.usersUC.UpdatePassword(
			r.Context(),
			user.Id,
			userUpdate.OldPassword,
			userUpdate.NewPassword,
			userUpdate.ConfirmPassword,
		)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err))
			return
		}

		render.Respond(w, r, responses.CreateSuccessResponse(mapModelResponse(updatedUser)))
	}
}

// LogoutAllAdmin ออกจากระบบทุกเซสชันของผู้ใช้ตาม ID (เฉพาะ superuser)
// LogoutAllAdmin logs out all sessions of a user by ID (superuser only)
// @Summary Logout all of user
// @Description Logout all session of user with id.
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User Id"
// @Success 200
// @Router /user/{id}/logoutall [get]
func (h *userHandler) LogoutAllAdmin() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		id, err := uuid.Parse(chi.URLParam(r, "id"))
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(httpErrors.ErrValidation(err)))
			return
		}

		err = h.usersUC.LogoutAll(ctx, id)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err))
			return
		}
	}
}

// ==================== NEW PUBLIC HANDLERS (ไม่ต้องใช้ JWT) ====================

// Register ลงทะเบียนผู้ใช้ใหม่ (สาธารณะ)
// Register creates a new user account (public)
// @Summary Register a new user
// @Description Create a new user account. No authentication required.
// @Tags auth
// @Accept json
// @Produce json
// @Param user body presenter.UserCreate true "User registration info"
// @Success 200 {object} responses.SuccessResponse[presenter.UserResponse]
// @Failure 400 {object} responses.ErrorResponse
// @Router /auth/register [post]

func (h *userHandler) Register() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req presenter.UserCreate
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err))
			return
		}
		if err := utils.ValidateStruct(r.Context(), req); err != nil {
			render.Render(w, r, responses.CreateErrorResponse(httpErrors.ErrValidation(err)))
			return
		}
		userModel := mapModel(&req)
		newUser, err := h.usersUC.CreateUser(r.Context(), userModel, req.ConfirmPassword)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err))
			return
		}
		render.Respond(w, r, responses.CreateSuccessResponse(mapModelResponse(newUser)))
	}
}

 
// Login เข้าสู่ระบบ (สาธารณะ)
// Login authenticates a user and returns tokens (public)
// @Summary Login user
// @Description Authenticate user with email and password.
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body presenter.UserSignIn true "Login credentials"
// @Success 200 {object} responses.SuccessResponse[presenter.Token]
// @Failure 401 {object} responses.ErrorResponse
// @Router /auth/login [post]
func (h *userHandler) Login() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req presenter.UserSignIn
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err))
			return
		}

		if err := utils.ValidateStruct(r.Context(), req); err != nil {
			render.Render(w, r, responses.CreateErrorResponse(httpErrors.ErrValidation(err)))
			return
		}

		// เรียก usecase SignIn / Call SignIn usecase
		accessToken, refreshToken, err := h.usersUC.SignIn(r.Context(), req.Email, req.Password)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err))
			return
		}

		render.Respond(w, r, responses.CreateSuccessResponse(presenter.Token{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			TokenType:    "Bearer",
		}))
	}
}

// RefreshToken สร้าง access token ใหม่โดยใช้ refresh token (สาธารณะ)
// RefreshToken generates a new access token using a valid refresh token
// @Summary Refresh access token
// @Description Get new access token using refresh token.
// @Tags auth
// @Accept json
// @Produce json
// @Param refresh body presenter.RefreshRequest true "Refresh token"
// @Success 200 {object} responses.SuccessResponse[presenter.Token]
// @Failure 401 {object} responses.ErrorResponse
// @Router /auth/refresh [post]
func (h *userHandler) RefreshToken() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req presenter.RefreshRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err))
			return
		}
		if req.RefreshToken == "" {
			render.Render(w, r, responses.CreateErrorResponse(httpErrors.ErrValidation(errors.New("refresh token required"))))
			return
		}
		accessToken, refreshToken, err := h.usersUC.Refresh(r.Context(), req.RefreshToken)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err))
			return
		}
		render.Respond(w, r, responses.CreateSuccessResponse(presenter.Token{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			TokenType:    "Bearer",
		}))
	}
}

// Logout ออกจากระบบ (สาธารณะ) - ใช้ refresh token เพื่อลบ session
// Logout invalidates the given refresh token (public)
// @Summary Logout user
// @Description Invalidate refresh token.
// @Tags auth
// @Accept json
// @Produce json
// @Param refresh body presenter.RefreshRequest true "Refresh token"
// @Success 200 {object} responses.SuccessResponse[string]
// @Router /auth/logout [post]
func (h *userHandler) Logout() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req presenter.RefreshRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err))
			return
		}

		if req.RefreshToken == "" {
			render.Render(w, r, responses.CreateErrorResponse(httpErrors.ErrValidation(errors.New("refresh token required"))))
			return
		}

		err := h.usersUC.Logout(r.Context(), req.RefreshToken)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err))
			return
		}

		render.Respond(w, r, responses.CreateSuccessResponse(map[string]string{"message": "logged out successfully"}))
	}
}

// ForgotPassword ขอรีเซ็ตรหัสผ่าน (สาธารณะ)
// ForgotPassword sends a password reset link to the user's email
// @Summary Forgot password
// @Description Send password reset link to email.
// @Tags auth
// @Accept json
// @Produce json
// @Param email body presenter.ForgotPassword true "User email"
// @Success 200 {object} responses.SuccessResponse[string]
// @Router /auth/forgot-password [post]
func (h *userHandler) ForgotPassword() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req presenter.ForgotPassword
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err))
			return
		}

		if err := utils.ValidateStruct(r.Context(), req); err != nil {
			render.Render(w, r, responses.CreateErrorResponse(httpErrors.ErrValidation(err)))
			return
		}

		// เรียก usecase ForgotPassword (ไม่คืนค่าอะไรนอกจาก error)
		err := h.usersUC.ForgotPassword(r.Context(), req.Email)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err))
			return
		}

		// สำเร็จเสมอ (ไม่เปิดเผยว่ามีอีเมลหรือไม่) / Always return success to prevent email enumeration
		render.Respond(w, r, responses.CreateSuccessResponse(map[string]string{"message": "reset link sent if email exists"}))
	}
}

// VerifyEmail ยืนยันอีเมลด้วย verification code (สาธารณะ)
// VerifyEmail verifies user email using the verification code
// @Summary Verify email
// @Description Verify user email with code from email link.
// @Tags auth
// @Accept json
// @Produce json
// @Param code query string true "Verification code"
// @Success 200 {object} responses.SuccessResponse[string]
// @Router /auth/verify-email [get]
func (h *userHandler) VerifyEmail() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		if code == "" {
			render.Render(w, r, responses.CreateErrorResponse(httpErrors.ErrValidation(errors.New("missing verification code"))))
			return
		}

		err := h.usersUC.Verify(r.Context(), code)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err))
			return
		}

		render.Respond(w, r, responses.CreateSuccessResponse(map[string]string{"message": "email verified successfully"}))
	}
}

// ResendVerification ส่งอีเมลยืนยันใหม่ (สาธารณะ)
// ResendVerification sends a new verification email to the user
// @Summary Resend verification email
// @Description Send a new verification link to user email.
// @Tags auth
// @Accept json
// @Produce json
// @Param email body presenter.ResendVerificationRequest true "User email"
// @Success 200 {object} responses.SuccessResponse[string]
// @Router /auth/resend-verification [post]
func (h *userHandler) ResendVerification() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req presenter.ResendVerificationRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err))
			return
		}

		if err := utils.ValidateStruct(r.Context(), req); err != nil {
			render.Render(w, r, responses.CreateErrorResponse(httpErrors.ErrValidation(err)))
			return
		}

		err := h.usersUC.ResendVerification(r.Context(), req.Email)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err))
			return
		}

		render.Respond(w, r, responses.CreateSuccessResponse(map[string]string{"message": "verification email sent"}))
	}
}

// ResetPassword รีเซ็ตรหัสผ่านโดยใช้ reset token (สาธารณะ)
// ResetPassword resets user password using reset token from email
// @Summary Reset password
// @Description Set new password using reset token.
// @Tags auth
// @Accept json
// @Produce json
// @Param code query string true "Reset token"
// @Param password body presenter.ResetPassword true "New password"
// @Success 200 {object} responses.SuccessResponse[string]
// @Router /auth/reset-password [post]
func (h *userHandler) ResetPassword() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.URL.Query().Get("code")
		if token == "" {
			render.Render(w, r, responses.CreateErrorResponse(httpErrors.ErrValidation(errors.New("missing reset token"))))
			return
		}

		var req presenter.ResetPassword
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err))
			return
		}

		if err := utils.ValidateStruct(r.Context(), req); err != nil {
			render.Render(w, r, responses.CreateErrorResponse(httpErrors.ErrValidation(err)))
			return
		}

		err := h.usersUC.ResetPassword(r.Context(), token, req.NewPassword, req.ConfirmPassword)
		if err != nil {
			render.Render(w, r, responses.CreateErrorResponse(err))
			return
		}

		render.Respond(w, r, responses.CreateSuccessResponse(map[string]string{"message": "password reset successfully"}))
	}
}

// ==================== Helper Functions ====================

// mapModel แปลง UserCreate presenter เป็น model User
// mapModel converts UserCreate presenter to User model
func mapModel(exp *presenter.UserCreate) *models.User {
	return &models.User{
		Name:        exp.Name,
		Email:       exp.Email,
		Password:    exp.Password,
		IsActive:    true,
		IsSuperUser: false,
	}
}

// mapModelResponse แปลง model User เป็น presenter UserResponse
// mapModelResponse converts User model to UserResponse presenter
func mapModelResponse(exp *models.User) *presenter.UserResponse {
	return &presenter.UserResponse{
		Id:          exp.Id,
		Name:        exp.Name,
		Email:       exp.Email,
		CreatedAt:   exp.CreatedAt,
		UpdatedAt:   exp.UpdatedAt,
		IsActive:    exp.IsActive,
		IsSuperUser: exp.IsSuperUser,
		Verified:    exp.Verified,
	}
}

// mapModelsResponse แปลง slice ของ model User เป็น slice ของ presenter UserResponse
// mapModelsResponse converts slice of User models to slice of UserResponse presenters
func mapModelsResponse(exp []*models.User) []*presenter.UserResponse {
	out := make([]*presenter.UserResponse, len(exp))
	for i, user := range exp {
		out[i] = mapModelResponse(user)
	}
	return out
}