package main

import (
	"fmt"
	"net/http"

	"github.com/thepralad/snet-college-networking/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)

	http.HandleFunc("/feeds", handlers.FeedsHandler)
	fmt.Println("Server starting at :8080")
	http.ListenAndServe(":8080", nil)
}
