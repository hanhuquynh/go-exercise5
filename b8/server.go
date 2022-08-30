package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	route.GET("/", homeHandler)

	route.GET("/:name", getParam)

	route.GET("/index", defautQuery)

	route.POST("/", postMethod)

	route.Run(":3000")
}

func homeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func getParam(c *gin.Context) {
	name := c.Param("name")

	c.JSON(http.StatusOK, gin.H{
		"message": "Name: " + name,
	})
}

func defautQuery(c *gin.Context) {
	addr := c.DefaultQuery("address", "Hà Nội")

	c.JSON(http.StatusOK, gin.H{
		"message": "Address: " + addr,
	})
}

func postMethod(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "POST Method",
	})
}
