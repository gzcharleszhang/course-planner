package main

import "net/http"

func rootHandler(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path

	_, err := w.Write([]byte(message))
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", rootHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
