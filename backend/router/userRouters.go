package router

import (
	"backend/db"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	crud := &db.UserCRUD{}
	var Login db.UserLogin
	if err := c.ShouldBindJSON(&Login); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	user, err := crud.GetUserByName(Login.UserName)
	if err != nil {
		c.JSON(404, gin.H{"error": "Invalid username or password"})
		return
	}

	if user.PassWord != Login.Password {
		c.JSON(404, gin.H{"error": "Invalid username or password"})
		return
	}

	c.JSON(200, gin.H{"message": "Login successful"})
}
