package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jparker15/Go-CRUD-API/initializers"
	"github.com/jparker15/Go-CRUD-API/models"
)

func PostsCreate(c *gin.Context) {

	//GET data from REQ bodyt
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//CREATE
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post) // pass pointer of data to Create

	if result.Error != nil {
		c.Status(400)
		return
	}

	//RETURN it
	c.JSON(200, gin.H{
		"post": post,
	})
}
