package routes

import (
	"net/http"

	"example.com/go-api-practice/models"
	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	user.ID = 1

	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not save user", "error": err.Error()})
	}
	c.JSON(http.StatusCreated, gin.H{"message": "user created successfully", "data": user})
}

func login(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	err = user.ValidateCredentials()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "logged in successfully"})
}
