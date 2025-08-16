package server

import (
	"ecom/config"
	"ecom/controller"
	"ecom/repo"
	"ecom/routes"
	"ecom/service"
	"github.com/gin-gonic/gin"
)

func StartServer(conf *config.Config) {
	r := gin.Default()
	routes.RegisterRoutes(r) // default

	userRepo := repo.NewUserRepo(conf.Ctx, conf.UserCollection)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	userController.RegisterRoutes(r)
	_ = r.Run(":" + conf.AppPort)
}
