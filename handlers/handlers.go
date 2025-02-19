package handlers

import(
	"net/http"
	"fmt"
	"github.com/thepralad/snet-college-networking/models"
	"github.com/thepralad/snet-college-networking/services"
	"github.com/thepralad/snet-college-networking/templates"

)



func HomeHandler(res http.ResponseWriter, req *http.Request){
	render.RenderTemplate(res, "index", "")
}

func RegisterHandler(res http.ResponseWriter, req *http.Request){
	req.ParseForm()
	email := req.FormValue("email")
	password := req.FormValue("password")

	message := "http://localhost:8080/add?email=" + email + "&password=" + password
  
	services.SendEmail([]string{email}, "Hey thanks for using SNET!", message)

	render.RenderTemplate(res, "index", "Please check you mailbox!")

}

func AddHandler(res http.ResponseWriter, req *http.Request){
	if req.Method != http.MethodGet{
		http.Error(res, "invalid method", http.StatusMethodNotAllowed)
		return
	}

	email := req.FormValue("email")
	password := req.FormValue("password")

	models.UpdateRow(email, password)

	// For testing purpose
	fmt.Fprintf(res, "%s entered timeline", email)

}
func LoginHandler(res http.ResponseWriter, req *http.Request){
	req.ParseForm()
	email := req.FormValue("email")
	password := req.FormValue("password")

	_ = password
	
	//For testing purpose
	fmt.Fprintf(res, "Logged in as : %v", email)
}