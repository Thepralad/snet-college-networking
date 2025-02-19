package handlers

import(
	"net/http"
	//"github.com/thepralad/snet-college-networking/models"
	"github.com/thepralad/snet-college-networking/services"
	"github.com/thepralad/snet-college-networking/templates"

)

func HomeHandler(res http.ResponseWriter, req *http.Request){
	render.RenderTemplate(res, "index", "")
}

//This handler sends an confirmation email
func RegisterPageHandler(res http.ResponseWriter, req *http.Request){
	req.ParseForm()
	email := req.FormValue("email")
	password := req.FormValue("password")

	message := "http://localhost:8080/add?email=" + email + "&password=" + password
  
	services.SendEmail([]string{email}, "Hey thanks for using SNET!", message)

	render.RenderTemplate(res, "index", "Please check your mailbox! Spam folder in particular.")

}


//This handler displays the login page
func LoginPageHandler(res http.ResponseWriter, req *http.Request){
	render.RenderTemplate(res, "login", "")
}