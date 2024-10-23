package main

import (
	"fiap/ancora/controller"
	"fiap/ancora/db"
	"fiap/ancora/helper"
	"fmt"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	port := "5000"
	path := "./html/"

	router := gin.Default()

	router.LoadHTMLGlob(path + "*.html")
	router.Static("/static", path+"static/")

	// Set up routes
	router.GET("/", controller.Index)

	// API
	router.GET("/api/v1/get", controller.IndexAPI)
	router.POST("/api/v1/post", controller.IndexAPI)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	err := router.Run(":" + port)
	fmt.Println(err)

}

func main() {
	helper.ConfigFile = "./conf/db.conf"
	if !helper.ValidateConfig() {
		fmt.Println("Configuração inválida")
		return
	}
	database := db.GetDB()
	err := database.Ping()
	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados")
		return
	}
	StartServer()
}
