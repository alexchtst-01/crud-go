package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// simulasi cpp dek apa apa harus jelas tipe datanya
type User struct {
	Name string `json:"name"`
}

// sama ini juga harus pake pointer
func GetUser(c *gin.Context) {
	payload := User{Name: "John Doe"}
	// kalo make gin.H
	// payload := gin.H{"name": "John Doe"}
	c.JSON(http.StatusOK, payload)
}

func CreateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// gin.H itu semacam constructor untuk buat json object field (map lah ini namanya)
	var payload = gin.H{"message": "User created", "user": user}

	c.JSON(http.StatusCreated, payload)
}
