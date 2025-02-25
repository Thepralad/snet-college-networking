package handlers

import (
	"fmt"
	"net/http"

	"github.com/thepralad/snet-college-networking/models"
	render "github.com/thepralad/snet-college-networking/templates"
	"github.com/thepralad/snet-college-networking/types"
	"github.com/thepralad/snet-college-networking/utils"
)

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	_, err := req.Cookie("session_token")
	if err != nil {
		http.Redirect(res, req, "/login", http.StatusFound)
		return
	}
	http.Redirect(res, req, "/feeds", http.StatusFound)
}

func RegisterHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {

		username := req.FormValue("username")
		email := req.FormValue("email")
		password := req.FormValue("password")
		deanery := req.FormValue("deanery")
		year := req.FormValue("year")

		if !utils.IsValidEmail(email) {
			render.RenderTemplate(res, "register", types.Message{Alert: "@salesiancollege.net email only"})
			return
		}

		err := models.InsertUser(username, email, password, deanery, year)
		if err != nil {
			fmt.Println(err)
			render.RenderTemplate(res, "register", types.Message{Alert: "Error registering user, please try again!"})
			return
		}

		userInfo := types.UserInfo{
			Bio:           req.FormValue("bio"),
			Gender:        req.FormValue("gender"),
			Phone:         req.FormValue("phone"),
			RelStatus:     req.FormValue("relationship"),
			TopArtist:     req.FormValue("top_artist"),
			LookingFor:    req.FormValue("looking_for"),
			InstaUsername: req.FormValue("instagram"),
			Email:         email,
		}

		models.InsertUserInfos(&userInfo)

		http.Redirect(res, req, "/login", http.StatusFound)
		return

	}
	render.RenderTemplate(res, "register", types.Message{Alert: ""})
}

func LoginHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		email := req.FormValue("email")
		password := req.FormValue("password")

		if !utils.IsValidEmail(email) {
			render.RenderTemplate(res, "login", types.Message{Alert: "Only @salesiancollege.net email accepted!"})
			return
		}

		_id, _, _password, _ := models.GetUserByEmail(email)
		if password != _password {
			render.RenderTemplate(res, "login", types.Message{Alert: "Incorrect email/password"})
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
			render.RenderTemplate(res, "login", types.Message{Alert: "Error creating session, please try again!"})
			return
		}

		http.Redirect(res, req, "/feeds", http.StatusFound)
		return
	}
	render.RenderTemplate(res, "login", types.Message{Alert: ""})
}

func LogoutHandler(res http.ResponseWriter, req *http.Request) {
	_, err := req.Cookie("session_token")
	if err != nil {
		http.Redirect(res, req, "/login", http.StatusFound)
		return
	}
	http.SetCookie(res, &http.Cookie{
		Name:   "session_token",
		Value:  "",
		MaxAge: 3600,
	})
	http.Redirect(res, req, "/login", http.StatusFound)
}
