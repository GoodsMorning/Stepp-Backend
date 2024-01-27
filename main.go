package main

import (
	"stepp-backend/initializer"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.InitEnv()
	initializer.ConnectDB()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	test := r.Group("/test");{
		test.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "from test",
			})
		})
	}
	r.Run()
}