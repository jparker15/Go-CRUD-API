package main

import (
	"github.com/jparker15/Go-CRUD-API/initializers"
	"github.com/jparker15/Go-CRUD-API/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
