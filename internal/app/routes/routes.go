package routes

import (
	"github.com/go-chi/chi"
	"github.com/gzcharleszhang/course-planner/internal/app/handlers/getTimelineHandler"
	"github.com/gzcharleszhang/course-planner/internal/app/handlers/getTimelinesHandler"
	"github.com/gzcharleszhang/course-planner/internal/app/handlers/getUserHandler"
	"github.com/gzcharleszhang/course-planner/internal/app/handlers/loginHandler"
	"github.com/gzcharleszhang/course-planner/internal/app/handlers/logoutHandler"
	"github.com/gzcharleszhang/course-planner/internal/app/handlers/newTimelineHandler"
	"github.com/gzcharleszhang/course-planner/internal/app/handlers/newUserHandler"
	"github.com/gzcharleszhang/course-planner/internal/app/middlewares"
)

func initUserRoutes(r chi.Router) {
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
			r.Get(getUserHandler.RouteURL, getUserHandler.Handler)
		})

		// admin routes
		r.Group(func(r chi.Router) {
			r.Use(middlewares.VerifyAdminMiddleware)
		})
	})
}

func initTimelineRoutes(r chi.Router) {
	r.Route("/timelines", func(r chi.Router) {
		// authenticated routes
		r.Group(func(r chi.Router) {
			r.Use(middlewares.VerifyAuthenticatedMiddleware)
			r.Get(getTimelinesHandler.RouteURL, getTimelinesHandler.Handler)
			r.Get(getTimelineHandler.RouteURL, getTimelineHandler.Handler)
			r.Post(newUserHandler.RouteURL, newTimelineHandler.Handler)
		})
	})
}

func InitRoutes(r chi.Router) {
	initUserRoutes(r)
	initTimelineRoutes(r)
}
