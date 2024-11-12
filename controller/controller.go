package controller

import (
	"fiap/ancora/db/dao"

	"github.com/gin-gonic/gin"
)

// Pages - Struct to hold the pages
func Login(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}

func Home(c *gin.Context) {
	username, _ := c.Cookie("username")
	getToken, _ := c.Cookie("token")
	if getToken == "" {
		c.Redirect(302, "/")
		return
	}
	var user = dao.GetUserByUsername(username)
	c.HTML(200, "home.html", gin.H{
		"user": user,
	})
}
