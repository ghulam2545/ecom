package repo

import (
	"context"
	"ecom/model"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepo struct {
	ctx            context.Context
	userCollection *mongo.Collection
}

func NewUserRepo(c context.Context, u *mongo.Collection) *UserRepo {
	return &UserRepo{userCollection: u}
}

func (u *UserRepo) List() ([]*model.User, error) {
	cursor, err := u.userCollection.Find(u.ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(u.ctx)

	var users []*model.User
	if err := cursor.All(u.ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}
