package handlers

import(
	"net/http"
	"fmt"
)

func InputMux(res http.ResponseWriter, req *http.Request){
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