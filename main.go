package main

import (
	"fmt"
	"net/http"

	"github.com/thepralad/snet-college-networking/handlers"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)

	http.HandleFunc("/logout", handlers.LogoutHandler)
	http.HandleFunc("/feeds", handlers.FeedsHandler)
	fmt.Println("Server starting at :8080")
	http.ListenAndServe(":8080", nil)
}
