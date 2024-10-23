package controller

import "github.com/gin-gonic/gin"

// Pages - Struct to hold the pages
func Index(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"title": "Main website",
	})
}
