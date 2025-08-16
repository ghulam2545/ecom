package repo

import (
	"context"
	"ecom/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repository interface {
	Create(p *model.Product) error
	GetByID(id string) (*model.Product, error)
	List() ([]*model.Product, error)
	Update(p *model.Product) error
	Delete(id string) error
}

type ProductRepo struct {
	productCollection *mongo.Collection
}

func NewProductRepo(p *mongo.Collection) *ProductRepo {
	return &ProductRepo{
		productCollection: p,
	}
}

func (r *ProductRepo) Create(p *model.Product) error {
	// TODO
	return nil
}

func (r *ProductRepo) GetByID(id string) (*model.Product, error) {
	// TODO
	return nil, nil
}

func (r *ProductRepo) List() ([]*model.Product, error) {
	cursor, err := r.productCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var products []*model.Product
	for cursor.Next(context.TODO()) {
		var product model.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (r *ProductRepo) Update(p *model.Product) error {
	// TODO
	return nil
}

func (r *ProductRepo) Delete(id string) error {
	// TODO
	return nil
}
