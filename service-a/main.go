package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/hello-a", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Service A!")
	})

	fmt.Println("Service A running on port 8081")

	time.Sleep(5 * time.Second)
	http.ListenAndServe(":8081", nil)
}
