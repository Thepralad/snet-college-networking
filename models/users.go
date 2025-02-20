package models

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


func InsertUser(email string, password string) error{
	db, err := sql.Open("mysql", "root:25802580@tcp(127.0.0.1:3306)/snet")	
	if err != nil{
		return err
	}
	defer db.Close()
	
	stmt, err := db.Prepare("INSERT INTO users(email, pass) VALUES(?, ?)")
	if err != nil{
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(email,password)
	if err != nil{
		return err
	}
	return nil
}
