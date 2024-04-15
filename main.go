package main

import (
	"fmt"

	"github.com/AnkitNayan83/gin-crud/controllers"
	"github.com/AnkitNayan83/gin-crud/initializers"
	"github.com/gin-gonic/gin"
)

// Init function runs before main function
func init() { 
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	fmt.Println("---------------------")
	r := gin.Default()
	r.POST("/posts", controllers.PostCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.DELETE("/posts/:id", controllers.PostDelete)
	r.Run() 
}