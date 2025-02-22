package handlers

import (
	"net/http"
	"strings"
	"github.com/thepralad/snet-college-networking/models"
	// "github.com/thepralad/snet-college-networking/services"
	"fmt"
	"github.com/thepralad/snet-college-networking/templates"
)

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	http.Redirect(res, req, "/login", http.StatusFound)
}

func RegisterHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost{

		username := req.FormValue("username")
		email := req.FormValue("email")
		password := req.FormValue("password")
		deanery := req.FormValue("deanery")
		year := req.FormValue("year")

		if !strings.HasSuffix(email, "@salesiancollege.net") {
			render.RenderTemplate(res, "register", "Only @salesiancollege.net email accepted!")
			return
		}

		err := models.InsertUser(username, email, password, deanery, year)		
		if err != nil {
		    render.RenderTemplate(res, "register", "Error registering user, please try again!")
		    return 
		}

		http.Redirect(res, req, "/login", http.StatusFound)
		return

	}
	render.RenderTemplate(res, "register", "")
}

func LoginHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost{
		email := req.FormValue("email")
		password := req.FormValue("password")

		if !strings.HasSuffix(email, "@salesiancollege.net") {
			render.RenderTemplate(res, "login", "Only @salesiancollege.net email accepted!")
			return
		}

		_username, _password, _:= models.GetUserByEmail(email)
		if password != _password{
			render.RenderTemplate(res, "login", "Incorrect email/password")
			return
		}

		fmt.Fprintln(res, "welcome: " + _username)
		return
	}
	render.RenderTemplate(res, "login", "")
}

