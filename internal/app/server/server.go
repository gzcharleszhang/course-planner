package server

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/gzcharleszhang/course-planner/internal/app/components/auth"
	"github.com/gzcharleszhang/course-planner/internal/app/middlewares"
	"github.com/gzcharleszhang/course-planner/internal/app/routes/userRoutes"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"time"
)

func StartServer(port int) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	// timeout in one minute
	r.Use(middleware.Timeout(60 * time.Second))
	// verify tokens
	r.Use(jwtauth.Verifier(auth.TokenAuth))
	// give request default permissions
	r.Use(middlewares.PermissionMiddleware)
	userRoutes.InitUserRoutes(r)
	fmt.Printf("Listening on port %v\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), r); err != nil {
		panic(err)
	}
}
