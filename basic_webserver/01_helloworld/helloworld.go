package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprint(w, "Hello World!")
		if err != nil {
			return
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error creating server: ", err)
		return
	}
}
