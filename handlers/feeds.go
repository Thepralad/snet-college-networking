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

	userId, err := models.GetUserIdFromSession(sessionToken.Value)
	if err != nil {
		http.Redirect(res, req, "/login", http.StatusFound)
		return
	}

	user, err := models.GetUserByUserId(userId)
	if err != nil {
		http.Redirect(res, req, "/login", http.StatusFound)
		return
	}
	feeds, err := models.GetFeedsArray()
	if err != nil {
		http.Redirect(res, req, "/login", http.StatusFound)
		return
	}

	userInfo, err := models.GetUserInfo(user.Email)
	if err != nil {
		fmt.Println(err)
	}
	render.RenderTemplate(res, "feeds", map[string]interface{}{
		"user":     user,
		"feeds":    feeds,
		"userInfo": userInfo,
	})
}

func PostFeedHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		sessionToken, err := req.Cookie("session_token")
		if err != nil {
			return
		}
		userId, err := models.GetUserIdFromSession(sessionToken.Value)
		if err != nil {
			return
		}
		models.PostFeed(userId, req.FormValue("content"))
		http.Redirect(res, req, "/feeds", http.StatusFound)
	}

}
