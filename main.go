package main

import (
	"fmt"
	"net/http"
)

func inputMux(res http.ResponseWriter, req *http.Request){
	if req.Method != http.MethodPost{
		fmt.Fprint(res, "invalid method")
		return
	}

	if req.FormValue("input") == "register"{
		fmt.Fprint(res, "you pressed register")
	}
	if req.FormValue("input") == "login"{
		fmt.Fprint(res, "you pressed login")
	}
}

func main(){
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/input", inputMux)
	fmt.Println("Server starting at :8080")
	http.ListenAndServe(":8080", nil)
}

