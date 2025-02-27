package models

import (
	"log"

	"github.com/thepralad/snet-college-networking/types"
)

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
