package server

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gzcharleszhang/course-planner/internal/app/components/users"
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
	// populate fields in context
	r.Use(contextMiddleware)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		panic(err)
	}
}

func contextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// missing user id
		if r.Header.Get("user_id") == "" {
			http.Error(w, "Error: missing user_id", 400)
			return
		}
		ctx := context.WithValue(r.Context(), "userId", users.UserId(r.Header.Get("user_id")))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
