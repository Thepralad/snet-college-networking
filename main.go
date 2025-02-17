package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/", landingPageHandler)
	http.ListenAndServe(":8080", nil)
}

func landingPageHandler(res http.ResponseWriter, req *http.Request){
	fmt.Fprint(res, "This is the landing page")
}