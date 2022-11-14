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

func PostsIndex(c *gin.Context) {
	//GET all posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	//Respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})

}
func PostsShow(c *gin.Context) {
	//GET id off URL
	id := c.Param("id")

	//GET all posts
	var post models.Post
	initializers.DB.Find(&post, id)

	//Respond with them
	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostsUpdate(c *gin.Context) {
	//Get the id off the url

	id := c.Param("id")

	//Get the data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//find the post to update
	var post models.Post
	initializers.DB.Find(&post, id)

	//update post
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}
