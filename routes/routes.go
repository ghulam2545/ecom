package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.Static("/js", "ui")
	router.StaticFile("/", "ui/index.html")
	router.StaticFile("/profile", "ui/profile.html")
}
