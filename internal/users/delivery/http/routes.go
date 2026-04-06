package http

import (
	"gorestapi/internal/middleware"
	"gorestapi/internal/users"

	"github.com/go-chi/chi/v5"
)

// MapUserRoute กำหนด routing ทั้งหมดของ module user
// MapUserRoute registers all user module routes
func MapUserRoute(router *chi.Mux, h users.Handlers, mw *middleware.MiddlewareManager) {
	// ==================== PUBLIC ROUTES (ไม่ต้องใช้ JWT) ====================
	// Public routes สำหรับ authentication และการจัดการบัญชี
	router.Route("/auth", func(r chi.Router) {
		// การลงทะเบียน / Registration
		r.Post("/register", h.Register())

		// การเข้าสู่ระบบ / Login
		r.Post("/login", h.Login())

		// รีเฟรช access token / Refresh access token
		r.Post("/refresh", h.RefreshToken())

		// ออกจากระบบ (ใช้ refresh token) / Logout (uses refresh token)
		r.Post("/logout", h.Logout())

		// ขอรีเซ็ตรหัสผ่าน (ส่งลิงก์ไปอีเมล) / Forgot password (send reset link)
		r.Post("/forgot-password", h.ForgotPassword())

		// ยืนยันอีเมล (GET พร้อม query param code) / Verify email (GET with code)
		r.Get("/verify-email", h.VerifyEmail())

		// ขอส่งอีเมลยืนยันใหม่ / Resend verification email
		r.Post("/resend-verification", h.ResendVerification())

		// รีเซ็ตรหัสผ่าน (POST พร้อม query param code และ body) / Reset password
		r.Post("/reset-password", h.ResetPassword())
	})

	// ==================== PROTECTED ROUTES (ต้องใช้ JWT) ====================
	// Protected routes - ต้องมี access token ที่ถูกต้อง
	router.Route("/user", func(r chi.Router) {
		// Middleware chain สำหรับทุก route ในกลุ่มนี้
		r.Use(mw.Verifier(true))      // ตรวจสอบและแยก access token
		r.Use(mw.Authenticator())     // ยืนยันความถูกต้องของ token
		r.Use(mw.CurrentUser())       // โหลดข้อมูลผู้ใช้จาก token ใส่ context
		r.Use(mw.ActiveUser())        // ตรวจสอบว่าผู้ใช้ยัง active อยู่

		// เส้นทางสำหรับผู้ใช้ปัจจุบัน (me) / Current user routes
		r.Get("/me", h.Me())
		r.Put("/me", h.UpdateMe())
		r.Patch("/me/updatepass", h.UpdatePasswordMe())

		// เส้นทางสำหรับ superuser เท่านั้น / Superuser-only routes
		r.Group(func(r chi.Router) {
			r.Use(mw.SuperUser()) // ตรวจสอบสิทธิ์ superuser

			// รายการผู้ใช้ (paginate) / List users
			r.Get("/", h.GetMulti())
			// สร้างผู้ใช้ (admin create) / Create user (admin)
			r.Post("/", h.Create())

			// เส้นทางที่มี parameter id / Routes with id parameter
			r.Route("/{id}", func(r chi.Router) {
				// ดึงข้อมูลผู้ใช้ตาม ID / Get user by ID
				r.Get("/", h.Get())
				// ลบผู้ใช้ / Delete user
				r.Delete("/", h.Delete())
				// อัปเดตข้อมูลผู้ใช้ / Update user
				r.Put("/", h.Update())
				// อัปเดตรหัสผ่านผู้ใช้ (admin) / Update user password (admin)
				r.Patch("/updatepass", h.UpdatePassword())
				// ออกจากระบบทุกเซสชันของผู้ใช้ / Logout all sessions of user
				r.Get("/logoutall", h.LogoutAllAdmin())
			})
		})
	})
}