package main

import (
	"UTokyoMeru/router"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Authorization, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())

	post := r.Group("/posts")
	{
		post.GET("", router.GetAllPosts)
		post.GET("/:PostId", router.GetAllReplies)
		post.POST("", router.CreatePost)
		post.POST("/:PostId", router.CreateReply)
	}

	err := r.Run(":8000")
	if err != nil {
		fmt.Println(err)
	}
}