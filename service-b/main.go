package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello-b", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Service B!")
	})

	fmt.Println("Service B running on port 8082")
	http.ListenAndServe(":8082", nil)
}
