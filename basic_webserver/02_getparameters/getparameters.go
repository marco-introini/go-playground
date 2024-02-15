package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "Guest"
		}
		fmt.Fprintf(w, "Hello, %s!", name)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error creating server: ", err)
		return
	}
}
