package server

import (
	"ecom/controller"
	"ecom/routes"
	"fmt"
	"github.com/gin-gonic/gin"
)

func StartServer(port string, userController *controller.UserController) {
	router := gin.Default()
	routes.RegisterRoutes(router, userController)

	err := router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}
}
