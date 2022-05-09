package main

import (
	"github.com/gin-gonic/gin"
	"github.com/aleqxan/Go-sql/db_client"
	"github.com/aleqxan/Go-sql/controllers"
)


func main() {
	db_client.InitializeDBConnection()

	r := gin.Default()

	r.POST(relativePath: "/", "controllers.CreatePost")
	r.GET(relativePath: "/", "controllers.GetPosts")
	r.GET(relativePath: "/:id", "controllers.GetPost")

	if err := r.Run(addr... ":5000"); err != nil {
		panic(err.Error())
	}
}