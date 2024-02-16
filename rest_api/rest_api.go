package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// curl localhost:8090/path/
	mux.HandleFunc("GET /path/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "got path\n")
	})

	//curl -X POST localhost:8090/path/
	mux.HandleFunc("POST /path/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "POST path\n")
	})

	// curl localhost:8090/task/f0cd2e/
	mux.HandleFunc("/task/{id}/", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "handling task with id=%v\n", id)
	})

	err := http.ListenAndServe("localhost:8090", mux)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
