package handlers

import (
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

		http.Redirect(res, req, "/login", http.StatusFound)
		return
	}

	render.RenderTemplate(res, "feeds", user)
}
