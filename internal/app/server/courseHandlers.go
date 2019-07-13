package server

import (
	"context"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path

	_, err := w.Write([]byte(message))
	if err != nil {
		panic(err)
	}
}

func HandleCourseRoutes() {
	http.HandleFunc("/courses", rootHandler)
	ctx := context.Background()

}
