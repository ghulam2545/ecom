package handler

import (
	"ecom/model"
	"ecom/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (c *UserHandler) RegisterRoutes(r *gin.Engine) {
	r.GET("/list", c.ListHandler)
	r.POST("/signup", c.SignupHandler)
}

func (c *UserHandler) ListHandler(ctx *gin.Context) {
	users, err := c.userService.List()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (c *UserHandler) SignupHandler(ctx *gin.Context) {
	var request model.SignupRequest
	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.userService.Signup(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
