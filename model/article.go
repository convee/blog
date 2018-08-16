package model

import (
	"github.com/convee/goboot/db/mysql"
)

type Article struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Createtime int64  `json:"create_time"`
}

func GetArticlesList(page int, perPage int) []Article {
	db := mysql.New("blog")
	var articles []Article
	offset := (page - 1) * perPage
	sql := "select id,title,content,create_time from article order by id desc limit ?, ?"
	rows, err := db.Query(sql, offset, perPage)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var article Article
		rows.Scan(&article.Id, &article.Title, &article.Content, &article.Createtime)
		articles = append(articles, article)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
	return articles
}
