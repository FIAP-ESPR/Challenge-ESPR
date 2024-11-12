package controller

import (
	"crypto/sha256"
	"fiap/ancora/db/dao"
	"fiap/ancora/model"

	"github.com/gin-gonic/gin"
)

func SignInPost(c *gin.Context) {
	username, _ := c.GetPostForm("username")
	password, _ := c.GetPostForm("password")
	var user = dao.GetUserByUsername(username)
	if user.Password == string(sha256.New().Sum([]byte(password))) {
		c.JSON(200, gin.H{
			"message": "Login successful",
		})
	}
	c.JSON(401, gin.H{
		"message": "Login failed",
	})
}

func SignUpPost(c *gin.Context) {
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
