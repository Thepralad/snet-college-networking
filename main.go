package main

import (
	"fmt"
	"net/http"
	"github.com/thepralad/snet-college-networking/handlers"
)



func main(){
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/input", handlers.InputMux)
	fmt.Println("Server starting at :8080")
	http.ListenAndServe(":8080", nil)
}

