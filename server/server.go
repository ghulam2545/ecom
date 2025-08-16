package server

import (
	"ecom/config"
	"ecom/handler"
	"ecom/repo"
	"ecom/routes"
	"ecom/service"
	"github.com/gin-gonic/gin"
)

func StartServer(conf *config.Config) {
	r := gin.Default()
	routes.RegisterRoutes(r) // default

	userRepo := repo.NewUserRepo(conf.UserCollection)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	productRepo := repo.NewProductRepo(conf.ProductCollection)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	userHandler.RegisterRoutes(r)
	productHandler.RegisterRoutes(r)
	_ = r.Run(":" + conf.AppPort)
}
