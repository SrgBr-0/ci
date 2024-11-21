// main.go
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	fmt.Fprintln(w, "Hello, World!")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
