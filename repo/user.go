package repo

import (
	"context"
	"ecom/model"
	"errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepo struct {
	userCollection *mongo.Collection
}

func NewUserRepo(u *mongo.Collection) *UserRepo {
	return &UserRepo{userCollection: u}
}

func (u *UserRepo) List(ctx context.Context) ([]*model.User, error) {
	cursor, err := u.userCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []*model.User
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserRepo) Save(ctx context.Context, user *model.User) (*model.User, error) {
	_, err := u.userCollection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepo) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User

	err := u.userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
