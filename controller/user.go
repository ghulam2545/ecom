package controller

import (
	"ecom/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) RegisterRoutes(r *gin.Engine) {
	r.GET("/list", c.ListController)
}

func (c *UserController) ListController(ctx *gin.Context) {
	users, err := c.userService.List()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}
