package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("")
	api.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping pong",
		})
	})
}
