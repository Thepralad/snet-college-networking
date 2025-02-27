package models

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/thepralad/snet-college-networking/types"
)

// DB is a global database connection pool
var DB *sql.DB

// InitDB initializes the database connection
func InitDB() error {
	godotenv.Load()
	var err error
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	connectionString := dbUsername + ":" + dbPassword + "@tcp(mysql-152cca74-snet.i.aivencloud.com:21979)/defaultdb"
	DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	err = DB.Ping()
	if err != nil {
		return err
	}

	return nil
}

// Inserts user to the database --- it takes (username, email, password, deanery, year)
// [used for registering]
func InsertUser(username string, email string, password string, deanery string, year string) error {
	stmt, err := DB.Prepare("INSERT INTO users(username, email, password, deanery, year) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, email, password, deanery, year)
	if err != nil {
		return err
	}

	return nil
}

// Get the user email --- it takes (email) and gives (userId, username, password)
// [user in logging in]
func GetUserByEmail(email string) (int, string, string, error) {
	stmt, err := DB.Prepare("SELECT id, username, password FROM users where email = ?")
	if err != nil {
		return 0, "", "", err
	}

	var userId int
	var username, password string
	err = stmt.QueryRow(email).Scan(&userId, &username, &password)
	if err != nil {
		return 0, "", "", err
	}
	return userId, username, password, nil
}

// Inserts a session token with its respective username -- it takes (userid, sessionToken)
func InsertSession(userID int, sessionToken string) error {
	_, err := DB.Exec("INSERT into session (user_id, session_token) VALUES (?, ?)", userID, sessionToken)
	if err != nil {
		return err
	}
	return nil
}

// Gets userId by sessionToken -- it takes (userId)
func GetUserIdFromSession(sessionToken string) (int, error) {
	var userId int
	err := DB.QueryRow("SELECT user_id from session where session_token = ?", sessionToken).Scan(&userId)
	if err != nil {
		return 0, err
	}
	return userId, nil
}

// Gets user basic info --- it takes (userId)
func GetUserByUserId(userId int) (*types.User, error) {
	var username, email, deanery, year, created_at string
	err := DB.QueryRow("SELECT username, email, deanery, year, created_at FROM users WHERE id = ?", userId).Scan(&username, &email, &deanery, &year, &created_at)
	if err != nil {
		return nil, err
	}
	return &types.User{
		Username: username,
		Email:    email,
		Deanery:  deanery,
		Year:     year,
		Created:  created_at,
	}, nil
}

func InsertUserInfos(userInfo *types.UserInfo) error {
	defaultDp := "https://media.istockphoto.com/id/1341046662/vector/picture-profile-icon-human-or-people-sign-and-symbol-for-template-design.jpg?s=612x612&w=0&k=20&c=A7z3OK0fElK3tFntKObma-3a7PyO8_2xxW0jtmjzT78="
	_, err := DB.Exec(`INSERT INTO user_info(bio, gender, phone, rel_status, top_artist,
					 looking_for, insta_username, email, img_url) values(?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		userInfo.Bio, userInfo.Gender, userInfo.Phone,
		userInfo.RelStatus, userInfo.TopArtist, userInfo.LookingFor, userInfo.InstaUsername, userInfo.Email, defaultDp)
	if err != nil {
		return err
	}
	return nil
}

func GetUserInfo(email string) (*types.UserInfo, error) {
	var userInfo types.UserInfo

	err := DB.QueryRow("SELECT bio, gender, phone, rel_status, top_artist, looking_for, insta_username, img_url FROM user_info WHERE email = ?", email).Scan(&userInfo.Bio, &userInfo.Gender, &userInfo.Phone, &userInfo.RelStatus, &userInfo.TopArtist, &userInfo.LookingFor, &userInfo.InstaUsername, &userInfo.Img_Url)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			// Return empty userInfo if no record found
			userInfo = types.UserInfo{
				Bio:           "",
				Gender:        "",
				Phone:         "",
				RelStatus:     "",
				TopArtist:     "",
				LookingFor:    "",
				InstaUsername: "",
			}
			return &userInfo, nil
		}
		return nil, err
	}
	return &userInfo, nil
}

func UpdateUserInfo(userInfo *types.UserInfo, email string) error {
	_, err := DB.Exec(`UPDATE user_info SET bio = ?, gender = ?, phone = ?, 
					rel_status = ?, top_artist = ?, looking_for = ?, insta_username = ?,
					img_url = ? 
					WHERE email = ?`,
		userInfo.Bio, userInfo.Gender, userInfo.Phone,
		userInfo.RelStatus, userInfo.TopArtist, userInfo.LookingFor,
		userInfo.InstaUsername, userInfo.Img_Url, email)
	if err != nil {
		return err
	}
	return nil
}
