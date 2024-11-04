package controller

import "github.com/gin-gonic/gin"

func IndexAPI(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}

func LoginValidate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}

func SignUp(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}
