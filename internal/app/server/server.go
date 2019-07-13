package server

import (
	"fmt"
	"net/http"
)

func StartServer(port int) {
	HandleCourseRoutes()
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		panic(err)
	}
}
