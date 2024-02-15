package main

import (
	"fmt"
	"net/http"
)

func main() {
	staticDir := http.Dir("./static")
	fileServer := http.FileServer(staticDir)
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error creating server: ", err)
		return
	}
}
