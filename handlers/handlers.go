package handlers

import (
	"fmt"
	//"github.com/thepralad/snet-college-networking/models"
	"github.com/thepralad/snet-college-networking/services"
	"github.com/thepralad/snet-college-networking/templates"
	"net/http"
)

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, "index", "")
}

// Sends an confirmation email
func RegisterPageHandler(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	email := req.FormValue("email")
	password := req.FormValue("password")

	url := "http://localhost:8080/adduser?email=" + email + "&password=" + password

	message := fmt.Sprintf(`<h3>Hello, User!</h3>
         <p>Thank you for signing up. Click 
         		<a href="%s" style="display:inline-block;">here</a>
          to verify your email.</p>`, url)

	services.SendEmail([]string{email}, "Hey thanks for using SNET!", message)

	render.RenderTemplate(res, "index", "Please check your mailbox! Spam folder in particular.")

}

// Adds a user to the database
func AddUserHandler(res http.ResponseWriter, req *http.Request) {
	// email := req.FormValue("email")
	// password := req.FormValue("password")

	// models.InsertUser(email, password)
	http.Redirect(res, req, "/login", http.StatusFound)

}

// Displays the login page
func LoginPageHandler(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, "login", "")
}