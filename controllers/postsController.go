package controllers

import (
	"github.com/AnkitNayan83/gin-crud/initializers"
	"github.com/AnkitNayan83/gin-crud/models"
	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {

	var body struct {
		Title string
		Body string
	}

	c.Bind(&body)

	if body.Title == "" {
		c.JSON(400, gin.H{
			"error": "Title is required",
		})
		return
	}

	if body.Body == "" {
		c.JSON(400, gin.H{
			"error": "Body is required",
		})
		return
	}

	post := models.Post{
		Title: body.Title,
		Body: body.Body,
	}

	result := initializers.DB.Create(&post) 

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": "failed to create post",
		})
		return
	}

	c.JSON(200,gin.H{
		"message":"post created",
		"post":post,
	})
}

func PostsIndex(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	if len(posts) == 0 {
		c.JSON(404, gin.H{
			"error": "No posts found",
		})
		return
	}

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {

	id := c.Param("id")

	if id == "" {
		c.JSON(400, gin.H{
			"error": "id is required",
		})
		return
	}

	var post models.Post
	initializers.DB.First(&post, id)

	if post == (models.Post{}) {
		c.JSON(404, gin.H{
			"error": "post not found",
		})
		return
	}

	c.JSON(200, post)
}

func PostUpdate(c *gin.Context) {

	id := c.Param("id")

	if id == "" {
		c.JSON(400, gin.H{
			"error": "id is required",
		})
		return
	}

	var post models.Post
	initializers.DB.First(&post, id)

	if post == (models.Post{}) {
		c.JSON(404, gin.H{
			"error": "post not found",
		})
		return
	}

	var body struct {
		Title string
		Body string
	}

	c.Bind(&body)

	if body.Title == "" {
		c.JSON(400, gin.H{
			"error": "Title is required",
		})
		return
	}

	if body.Body == "" {
		c.JSON(400, gin.H{
			"error": "Body is required",
		})
		return
	}

	post.Title = body.Title
	post.Body = body.Body

	initializers.DB.Save(&post)

	c.JSON(200, post)
}

func PostDelete(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(400, gin.H{
			"error": "id is required",
		})
		return
	}

	var post models.Post

	initializers.DB.First(&post, id)

	if post == (models.Post{}) {
		c.JSON(404, gin.H{
			"error": "post not found",
		})
		return
	}

	initializers.DB.Delete(&post)

	c.JSON(200, gin.H{
		"message": "post deleted",
	})
}