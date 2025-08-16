package service

import (
	"ecom/model"
	"ecom/repo"
	"github.com/gin-gonic/gin"
)

type Service interface {
	CreateProduct(name, description string, price float64, stock int) (*model.Product, error)
	GetProduct(id string) (*model.Product, error)
	ListProducts(ctx *gin.Context) ([]*model.Product, error)
	UpdateProduct(p *model.Product) error
	DeleteProduct(id string) error
}

type ProductService struct {
	productRepo *repo.ProductRepo
}

func NewProductService(productRepo *repo.ProductRepo) *ProductService {
	return &ProductService{productRepo: productRepo}
}

func (p *ProductService) CreateProduct(name, description string, price float64, stock int) (*model.Product, error) {
	// TODO: implement product creation logic (call repo)
	return nil, nil
}

func (p *ProductService) GetProduct(id string) (*model.Product, error) {
	// TODO: implement fetch by ID logic (call repo)
	return nil, nil
}

func (p *ProductService) ListProducts(ctx *gin.Context) ([]*model.Product, error) {
	return p.productRepo.List()
}

func (p *ProductService) UpdateProduct(product *model.Product) error {
	// TODO: implement update logic (call repo)
	return nil
}

func (p *ProductService) DeleteProduct(id string) error {
	// TODO: implement delete logic (call repo)
	return nil
}
