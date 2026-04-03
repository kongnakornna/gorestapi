package http

import (
	"gorestapi/internal/auth"
	"gorestapi/internal/middleware"

	"github.com/go-chi/chi/v5"
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
