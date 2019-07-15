package userRoutes

import (
	"github.com/go-chi/chi"
	"github.com/gzcharleszhang/course-planner/internal/app/handlers/userHandlers"
)

func InitUserRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/", userHandlers.NewUserHandler)
	})
}
