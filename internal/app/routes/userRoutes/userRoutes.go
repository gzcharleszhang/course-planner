package userRoutes

import (
	"github.com/go-chi/chi"
	"github.com/gzcharleszhang/course-planner/internal/app/handlers/userHandlers"
	"github.com/gzcharleszhang/course-planner/internal/app/middlewares"
)

func InitUserRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		// unauthenticated routes
		r.Group(func(r chi.Router) {
			r.Post("/register", userHandlers.NewUserHandler)
			r.Post("/login", userHandlers.LoginHandler)
		})

		// authenticated routes
		r.Group(func(r chi.Router) {
			r.Use(middlewares.VerifyAuthenticatedMiddleware)
			r.Post("/logout", userHandlers.LogoutHandler)
		})

		// admin routes
		r.Group(func(r chi.Router) {
			r.Use(middlewares.VerifyAdminMiddleware)
		})
	})
}
