package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/kongnakornna/gorestapi/internal/auth"
	"github.com/kongnakornna/gorestapi/internal/middleware"
)

func MapAuthRoute(router *chi.Mux, h auth.Handlers, mw *middleware.MiddlewareManager) {
	// Auth routes
	router.Route("/auth", func(r chi.Router) {
		// Public routes
		r.Group(func(r chi.Router) {
			r.Post("/login", h.SignIn())
			r.Get("/publickey", h.GetPublicKey())
			r.Get("/verifyemail", h.VerifyEmail())
			r.Post("/forgotpassword", h.ForgotPassword())
			r.Patch("/resetpassword", h.ResetPassword())
		})
		r.Group(func(r chi.Router) {
			r.Use(mw.Verifier(false))
			r.Use(mw.Authenticator())
			r.Get("/refresh", h.RefreshToken())
			r.Get("/logout", h.Logout())
			r.Get("/logoutall", h.LogoutAllToken())
		})
	})
}
