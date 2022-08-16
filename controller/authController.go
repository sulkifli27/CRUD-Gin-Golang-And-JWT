package controller

import (
	"fmt"
	"golang-test/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}
	u.Username = input.Username
	u.Password = input.Password
	token, err := models.LoginCheck(u.Username, u.Password, db)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "login success", "token": token})
}


func Register(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    var input RegisterInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    u := models.User{}
    u.Username = input.Username
    u.Password = input.Password

    _, err := u.SaveUser(db)

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user := map[string]string{
        "username": input.Username,
    }

    c.JSON(http.StatusOK, gin.H{"message": "registration success", "data": user})

}