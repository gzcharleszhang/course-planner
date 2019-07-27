package userRoutes

import (
	"github.com/go-chi/chi"
	"github.com/gzcharleszhang/course-planner/internal/app/handlers/loginHandler"
	"github.com/gzcharleszhang/course-planner/internal/app/handlers/logoutHandler"
	"github.com/gzcharleszhang/course-planner/internal/app/handlers/newUserHandler"
	"github.com/gzcharleszhang/course-planner/internal/app/middlewares"
)

func InitUserRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		// unauthenticated routes
		r.Group(func(r chi.Router) {
			r.Post(newUserHandler.RouteURL, newUserHandler.Handler)
			r.Post(loginHandler.RouteURL, loginHandler.Handler)
		})

		// authenticated routes
		r.Group(func(r chi.Router) {
			r.Use(middlewares.VerifyAuthenticatedMiddleware)
			r.Post(logoutHandler.RouteURL, logoutHandler.Handler)
		})

		// admin routes
		r.Group(func(r chi.Router) {
			r.Use(middlewares.VerifyAdminMiddleware)
		})
	})
}
