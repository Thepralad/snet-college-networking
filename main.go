package main

import (
	"fmt"
	"net/http"
	"github.com/thepralad/snet-college-networking/handlers"
)

func main(){
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/register", handlers.RegisterPageHandler)
	http.HandleFunc("/login", handlers.LoginPageHandler)
	http.HandleFunc("/adduser", handlers.AddUserHandler)

	fmt.Println("Server starting at :8080")
	http.ListenAndServe(":8080", nil)
}

