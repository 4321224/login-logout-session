package main

import (
	"fmt"
	"net/http"

	"login-logout-session/go-auth/controllers"
)

func main() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/login", controllers.Login)
	fmt.Println("Server is running on port 4000")
	http.ListenAndServe(":4000", nil)
}
