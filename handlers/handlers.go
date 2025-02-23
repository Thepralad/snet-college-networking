package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/thepralad/snet-college-networking/models"
	render "github.com/thepralad/snet-college-networking/templates"
	"github.com/thepralad/snet-college-networking/utils"
)

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	sessionToken, err := req.Cookie("session_token")
	if err != nil {
		http.Redirect(res, req, "/login", http.StatusFound)
		return
	}

	userId, err := models.GetSession(sessionToken.Value)
	if err != nil {
		http.Redirect(res, req, "/login", http.StatusFound)
		return
	}

	username, err := models.GetUsernameByUserId(userId)
	if err != nil {
		log.Printf("Error getting username: %v", err)
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "text/html")
	res.Write([]byte(fmt.Sprintf("Welcome: %s", username)))
}

func RegisterHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {

		username := req.FormValue("username")
		email := req.FormValue("email")
		password := req.FormValue("password")
		deanery := req.FormValue("deanery")
		year := req.FormValue("year")

		if !utils.IsValidEmail(email) {
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
	if req.Method == http.MethodPost {
		email := req.FormValue("email")
		password := req.FormValue("password")

		if !utils.IsValidEmail(email) {
			render.RenderTemplate(res, "register", "Only @salesiancollege.net email accepted!")
			return
		}

		_id, _username, _password, _ := models.GetUserByEmail(email)
		if password != _password {
			render.RenderTemplate(res, "login", "Incorrect email/password")
			return
		}

		sessionToken := utils.GenerateSessionToken()

		http.SetCookie(res, &http.Cookie{
			Name:   "session_token",
			Value:  sessionToken,
			MaxAge: 3600,
		})

		err := models.InsertSession(_id, sessionToken)
		if err != nil {
			log.Println("Error inserting session:", err)
			render.RenderTemplate(res, "login", "Error creating session, please try again!")
			return
		}

		fmt.Fprintln(res, "welcome: "+_username)

		return
	}
	render.RenderTemplate(res, "login", "")
}
