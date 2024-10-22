package main

import (
	"fiap/ancora/controller"
	"fmt"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	port := "8080"
	path := "./html/"

	router := gin.Default()

	router.LoadHTMLGlob(path + "*.html")
	router.Static("/static", path+"static/")

	// Set up routes
	router.GET("/", controller.Index)

	// API
	router.POST("/api/v1", controller.IndexAPI)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	err := router.Run(":" + port)
	fmt.Println(err)

}

func main() {
	StartServer()
}
