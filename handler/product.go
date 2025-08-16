package handler

import (
	"ecom/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductHandler struct {
	productService *service.ProductService
}

func NewProductHandler(productService *service.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (c *ProductHandler) RegisterRoutes(r *gin.Engine) {
	r.GET("/products", c.ListProductsHandler)
}

func (c *ProductHandler) ListProductsHandler(ctx *gin.Context) {
	products, err := c.productService.ListProducts(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}
