package routes

import (
	"ecom/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, userController *controller.UserController) {
	api := router.Group("")
	api.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping pong",
		})
	})

	api.GET("/list", userController.ListController)
}
