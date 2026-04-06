package users

import "net/http"

// Handlers interface สำหรับ HTTP handlers ทั้งหมด
// Handlers interface defines all HTTP handlers
type Handlers interface {
	// Protected / Admin handlers
	Create() func(w http.ResponseWriter, r *http.Request)
	Get() func(w http.ResponseWriter, r *http.Request)
	GetMulti() func(w http.ResponseWriter, r *http.Request)
	Delete() func(w http.ResponseWriter, r *http.Request)
	Update() func(w http.ResponseWriter, r *http.Request)
	Me() func(w http.ResponseWriter, r *http.Request)
	UpdateMe() func(w http.ResponseWriter, r *http.Request)
	UpdatePassword() func(w http.ResponseWriter, r *http.Request)
	UpdatePasswordMe() func(w http.ResponseWriter, r *http.Request)
	LogoutAllAdmin() func(w http.ResponseWriter, r *http.Request)

	// Public handlers (ไม่ต้องใช้ JWT)
	Register() func(w http.ResponseWriter, r *http.Request)
	Login() func(w http.ResponseWriter, r *http.Request)
	RefreshToken() func(w http.ResponseWriter, r *http.Request)
	Logout() func(w http.ResponseWriter, r *http.Request)
	ForgotPassword() func(w http.ResponseWriter, r *http.Request)
	VerifyEmail() func(w http.ResponseWriter, r *http.Request)
	ResendVerification() func(w http.ResponseWriter, r *http.Request)
	ResetPassword() func(w http.ResponseWriter, r *http.Request)
}