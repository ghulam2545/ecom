package server

import (
	"ecom/routes"
	"fmt"
	"github.com/gin-gonic/gin"
)

func StartServer(port string) {
	router := gin.Default()
	routes.RegisterRoutes(router)

	err := router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}
}
