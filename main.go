package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ini perihal dummy data
type book struct {
	ID       string `json:"id"`
	Titile   string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

// dummy data
var books = []book{
	{ID: "1", Titile: "loosing time", Author: "author 1", Quantity: 12},
	{ID: "2", Titile: "paying time", Author: "author 2", Quantity: 21},
	{ID: "3", Titile: "gain time", Author: "author 3", Quantity: 11},
	{ID: "4", Titile: "hoping in time", Author: "author 1", Quantity: 10},
}

// kalo di javascript ini namanya controller
func getBook(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// ini cerintanya index.js (entry point access servernya disini)
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/getbook", getBook)
	r.Run()
	// listen and serve on 0.0.0.0:8080
}
