package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/thepralad/snet-college-networking/types"
)

// DB is a global database connection pool
var DB *sql.DB

// InitDB initializes the database connection
func initDB() error {
	var err error
	DB, err = sql.Open("mysql", "root:25802580@tcp(127.0.0.1:3306)/snet")
	if err != nil {
		return err
	}
	err = DB.Ping()
	if err != nil {
		return err
	}

	return nil
}

func InsertUser(username string, email string, password string, deanery string, year string) error {
	initDB()
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

func GetUserByEmail(email string) (int, string, string, error) {
	initDB()
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

func InsertSession(userID int, sessionToken string) error {
	initDB()
	_, err := DB.Exec("INSERT into session (user_id, session_token) VALUES (?, ?)", userID, sessionToken)
	if err != nil {
		return err
	}
	return nil
}

func GetSession(sessionToken string) (int, error) {
	initDB()
	var userId int
	err := DB.QueryRow("SELECT user_id from session where session_token = ?", sessionToken).Scan(&userId)
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func GetUserByUserId(userId int) (*types.User, error) {
	initDB()
	var username, email, deanery, year string
	err := DB.QueryRow("SELECT username, email, deanery, year FROM users WHERE id = ?", userId).Scan(&username, &email, &deanery, &year)
	if err != nil {
		return nil, err
	}
	return &types.User{
		Username: username,
		Email:    email,
		Deanery:  deanery,
		Year:     year,
	}, nil
}
