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
	"os"
	"time"
)

func StartServer(port int) {
	err := LoadEnv()
	if err != nil {
		log.Fatalf("Error: failed to load environment variables %v", err)
	}
	r := SetupRouter()
	fmt.Printf("Listening on port %v\n", port)
	y, m, d := time.Now().Date()
	// create file for error logging
	errorLog, err := os.OpenFile(fmt.Sprintf("%v-%v-%v_err.log", y, m, d), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	errLogger := log.New(errorLog, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), r); err != nil {
		// log error
		errLogger.Printf("%v", err)
		// print to stderr
		log.Printf("%v", err)
	}
	fmt.Print("hello")
}

func LoadEnv() error {
	return godotenv.Load()
}

func SetupRouter() *chi.Mux {
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
	return r
}
