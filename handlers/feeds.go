package handlers

import (
	"fmt"
	"net/http"

	"github.com/thepralad/snet-college-networking/models"
	render "github.com/thepralad/snet-college-networking/templates"
)

func FeedsHandler(res http.ResponseWriter, req *http.Request) {
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

	user, err := models.GetUserByUserId(userId)
	if err != nil {
		render.RenderTemplate(res, "login", "Error creating session, please try again!")
	}

	fmt.Fprintf(res, "Username: %s, Email: %v, Deanery: %v", user.Username, user.Email, user.Deanery)
}
