package controller

import "github.com/gin-gonic/gin"

// Pages - Struct to hold the pages
func Login(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{
		"title": "Main website",
	})
}
