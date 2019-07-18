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
			r.Post("/", userHandlers.NewUserHandler)
			r.Get("/", userHandlers.LoginHandler)
		})

		// authenticated routes
		r.Group(func(r chi.Router) {
			r.Use(middlewares.VerifyAuthenticatedMiddleware)
		})

		// admin routes
		r.Group(func(r chi.Router) {
			r.Use(middlewares.VerifyAdminMiddleware)
		})
	})
}
