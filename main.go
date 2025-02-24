package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/thepralad/snet-college-networking/handlers"
	"github.com/thepralad/snet-college-networking/models"
)

func main() {
	if err := models.InitDB(); err != nil {
		log.Fatal(err)
	}
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)

	http.HandleFunc("/logout", handlers.LogoutHandler)
	http.HandleFunc("/feeds", handlers.FeedsHandler)

	http.HandleFunc("/post", handlers.PostFeedHandler)
	fmt.Println("Server starting at :8080")
	http.ListenAndServe(":8080", nil)
}
