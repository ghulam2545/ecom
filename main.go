package main

import (
	"ecom/configuration"
	"ecom/controller"
	"ecom/repo"
	"ecom/server"
	"ecom/service"
)

func main() {
	conf := configuration.Configurations()
	port := conf.AppPort
	ctx := conf.Ctx
	userColl := conf.UserCollection

	userRepo := repo.NewUserRepo(ctx, userColl)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	server.StartServer(port, userController)
}
