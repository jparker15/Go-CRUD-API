package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jparker15/Go-CRUD-API/controllers"
	"github.com/jparker15/Go-CRUD-API/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.Run() // listen and serve on 0.0.0.0:8080
}
