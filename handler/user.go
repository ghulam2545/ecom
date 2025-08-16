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
	//r.GET("/list", c.ListHandler)
	r.POST("/signup", c.SignupHandler)
	r.POST("/login", c.LoginHandler)

	api := r.Group("")
	api.Use(service.JWTAuthMiddleware())
	{
		api.GET("/admin", service.RequireRoles("ADMIN"), func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "only admins can see this"})
		})

		api.GET("/user", service.RequireAnyRole("ADMIN", "USER"), func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "user or admin access"})
		})
	}
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

func (c *UserHandler) LoginHandler(ctx *gin.Context) {
	var request model.LoginRequest
	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, _, err := c.userService.Login(&request)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
