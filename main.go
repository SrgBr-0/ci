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
	http.HandleFunc("/", handler) // Устанавливаем обработчик для корня
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil) // Запускаем сервер на порту 8080
}
