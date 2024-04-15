package main

import (
	"github.com/AnkitNayan83/gin-crud/initializers"
	"github.com/AnkitNayan83/gin-crud/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}