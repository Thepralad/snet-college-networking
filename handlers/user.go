package handlers

import (
	"fmt"
	"net/http"

	"github.com/thepralad/snet-college-networking/models"
	"github.com/thepralad/snet-college-networking/services"
	render "github.com/thepralad/snet-college-networking/templates"
	"github.com/thepralad/snet-college-networking/types"
)

func EditProfileHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
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
			return
		}

		imgFile, _, err := req.FormFile("image")
		imgUrl, _ := services.UploadImg(imgFile)
		userInfo := types.UserInfo{
			Bio:           req.FormValue("bio"),
			Gender:        req.FormValue("gender"),
			Phone:         req.FormValue("phone"),
			RelStatus:     req.FormValue("relationship"),
			TopArtist:     req.FormValue("top_artist"),
			LookingFor:    req.FormValue("looking_for"),
			InstaUsername: req.FormValue("instagram"),
			Img_Url:       imgUrl,
		}

		if imgUrl == "" {
			existingUserInfo, _ := models.GetUserInfo(user.Email)
			userInfo.Img_Url = existingUserInfo.Img_Url
		}

		err = models.UpdateUserInfo(&userInfo, user.Email)
		if err != nil {
			fmt.Print(err)
		}
		http.Redirect(res, req, "/feeds", http.StatusFound)
	}

	if req.Method == http.MethodGet {
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
		userInfo, err := models.GetUserInfo(user.Email)
		if err != nil {
			fmt.Println(err)
		}
		err = render.RenderTemplate(res, "edit", map[string]interface{}{
			"user":     user,
			"userInfo": userInfo,
		})
		if err != nil {
			fmt.Fprintln(res, err)
		}
		return
	}

}
