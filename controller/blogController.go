package controller

import (
	"golang-test/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type BlogInput struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

func GetAllBlog(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var data []models.Blog
	db.Find(&data)
    c.JSON(http.StatusOK, gin.H{"code" : "200", "message": "success get data", "data": data })
}

func GetBlogById(c *gin.Context) { 
    // get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	var data models.Blog
	if err := db.Where("id = ?", c.Param("id")).First(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data not found!"})
		return
	}
    c.JSON(http.StatusOK, gin.H{"code" : "200", "message": "success get data", "data": data })
}

func CreateBlog(c *gin.Context) {
	// Validate input
	var input BlogInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

   // get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	// Create Blog
    slug := slug.Make(input.Title)
	data := models.Blog{Title: input.Title, Body: input.Body, Slug: slug}
	db.Create(&data)

	c.JSON(http.StatusOK, gin.H{"code": "200", "message": "success create data", "data": data})
}

func UpdateBlog(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var data models.Blog
	if err := db.Where("id = ?", c.Param("id")).First(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data not found!"})
		return
	}

	// Validate input
	var input BlogInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Blog
    slug := slug.Make(input.Title)
	updatedInput.Title = input.Title
	updatedInput.Body = input.Body
    updatedInput.Slug = slug
	updatedInput.UpdatedAt = time.Now()

	db.Model(&data).Updates(updatedInput)

    c.JSON(http.StatusOK, gin.H{"code" : "200", "message": "success update data", "data": data })

}

func DeleteBlog(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var data models.Blog
	if err := db.Where("id = ?", c.Param("id")).First(&data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Blog not found!"})
		return
	}
	db.Delete(&data)

	c.JSON(http.StatusOK, gin.H{"code" : "200", "message": "success delete data"})
}

