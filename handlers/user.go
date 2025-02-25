package handlers

import (
	"fmt"
	"net/http"

	"github.com/thepralad/snet-college-networking/models"
	render "github.com/thepralad/snet-college-networking/templates"
)

func EditProfileHandler(res http.ResponseWriter, req *http.Request) {
	sessionToken, err := req.Cookie("session_token")
	if err != nil {
		return
	}
	userId, err := models.GetUserIdFromSession(sessionToken.Value)
	if err != nil {
		return
	}

	user, err := models.GetUserByUserId(userId)
	if err != nil {
		http.Redirect(res, req, "/login", http.StatusFound)
		return
	}
	err = render.RenderTemplate(res, "edit", user)
	if err != nil {
		fmt.Fprintln(res, err)
	}
}
