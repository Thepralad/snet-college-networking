package models

import (
	"fmt"
	"log"

	"github.com/thepralad/snet-college-networking/types"
)

func PokeUser(from_id int, to_id int) error {
	exist, err := PokeExist(from_id, to_id)
	if err != nil {
		return err
	}

	if exist {
		fmt.Println("The poke already exists")
		return nil
	}
	_, err = DB.Exec("INSERT INTO pokes(from_id, to_id) VALUES(?,?)", from_id, to_id)
	if err != nil {
		return err
	}
	return nil
}

func PokeExist(from_id int, to_id int) (bool, error) {
	var exists bool
	err := DB.QueryRow("SELECT EXISTS(SELECT 1 FROM pokes WHERE from_id = ? AND to_id = ?)", from_id, to_id).Scan(&exists)
	if err != nil {
		return exists, err
	}
	return exists, nil
}
func GetPokes(userId int) ([]types.Poke, error) {
	var pokesFrom []types.Poke
	rows, err := DB.Query("SELECT from_id FROM pokes where to_id = ?", userId)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for rows.Next() {
		var from_id int
		err := rows.Scan(&from_id)
		if err != nil {
			return nil, err
		}

		var poke types.Poke
		row := DB.QueryRow("SELECT username, users.email, img_url FROM users JOIN user_info ON users.email = user_info.email WHERE users.id = ?", from_id)
		err = row.Scan(&poke.Username, &poke.Email, &poke.Img_Url)
		if err != nil {
			return nil, err
		}
		pokesFrom = append(pokesFrom, poke)
	}
	return pokesFrom, nil
}
