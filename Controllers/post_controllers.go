package controllers

import (
	"net/http"
	"time"
	"Strconv"
	"github.com/gin-gonic/gin"
	"github.com/aleqxan/Go-sql/db_client"
)

type Post struct {
	ID			int64 `json:"id"`
	Title		string `json:"title"`
	Content		string `json:"content"`
	CreatedAt 	time.Time `json: "created_at" db:"created_at"`
}

func CreatePost(c *gin.Context){
	var reqBody Post
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": true,
			"message": "Invalid request body",
		})
		return
	}

	res, _ := db_client.DBClient.Exec(query: "INSERT INTO posts (title, content) VALUES (?, ?);",
	reqBody.Title,
	reqBody.Content,
	)

	res.LastInsertId()
	res.RowsAffected()

	id, _ := res.LastInsertId()
	
	c.JSON(http.StatusCreated, gin.H{
		"error": false,
		"id": id,
	})
}

func GetPosts(c *gin.Context) {
	var posts []Post

	db_client.DBClient.Select(&Posts, query: "SELECT id, title, content, created_at FROM posts")

	c.JSON(http.StatusOK, posts)
}

func GetPost(c *gin.Context) {
	idStr := c.Param(key: "id")
	id, _ := strconv.Atoi(idStr)
	var post Post
	
	db_client.DBClient.Get(&post, query: "SELECT id, title, content, created_at FROM posts WHERE id = ?;", id)
	
	if err != nil{
		fmt.Println(err.Error)
	}

	c.JSON(http.StatusOK, post)
}