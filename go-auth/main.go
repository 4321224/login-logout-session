package main

import (
	"fmt"
	"net/http"

	"login-logout-session/go-auth/controllers"
)

func main() {
	http.HandleFunc("/", controllers.Index)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
