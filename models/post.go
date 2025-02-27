package models

import "github.com/thepralad/snet-college-networking/types"

func PostFeed(userId int, content string) error {
	stmt, err := DB.Prepare("INSERT INTO feeds(user_id, content) Values(?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userId, content)
	if err != nil {
		return err
	}

	return nil
}

func GetFeedsArray() ([]types.Post, error) {
	var postArr []types.Post
	rows, err := DB.Query(`SELECT users.username, users.email, user_info.img_url, feeds.content, feeds.created_at, feeds.metric
			FROM feeds
			JOIN users ON feeds.user_id = users.id
			JOIN user_info ON user_info.email = users.email`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var post types.Post
		_ = rows.Scan(&post.Username, &post.Email, &post.Img_Url, &post.Content, &post.Created, &post.Metric)
		postArr = append(postArr, post)
	}
	return postArr, nil
}
