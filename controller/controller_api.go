package controller

import (
	"crypto/sha256"
	"fiap/ancora/db/dao"
	"fiap/ancora/model"

	"github.com/gin-gonic/gin"
)

func SignInPost(c *gin.Context) {
	/* change domain before launch to production */
	domain := "localhost"
	if c.Request.Method != "GET" {
		c.JSON(405, gin.H{
			"message": "Method not allowed",
		})
	}
	username, _ := c.GetPostForm("username")
	password, _ := c.GetPostForm("password")
	var user = dao.GetUserByUsername(username)
	if user.Password == string(sha256.New().Sum([]byte(password))) {
		c.JSON(200, gin.H{
			"message": "Login successful",
		})
		token := string(sha256.New().Sum([]byte(username)))
		c.SetCookie("token", token, 86400, "/", domain, false, true)
		c.SetCookie("username", user.Username, 86400, "/", domain, false, true)
		c.Redirect(302, "/home")
		return
	}
	c.JSON(401, gin.H{
		"message": "Login failed",
	})
}

func SignUpPost(c *gin.Context) {
	if c.Request.Method != "POST" {
		c.JSON(405, gin.H{
			"message": "Method not allowed",
		})
	}
	username, _ := c.GetPostForm("username")
	password, _ := c.GetPostForm("password")
	user := model.User{
		Username: username,
		Password: password,
	}
	dao.InsertUser(user)
	c.JSON(200, gin.H{
		"message": "User created",
	})
}
