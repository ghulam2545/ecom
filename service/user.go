package service

import (
	"ecom/model"
	"ecom/repo"
)

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService(u *repo.UserRepo) *UserService {
	return &UserService{userRepo: u}
}

func (u *UserService) List() ([]*model.User, error) {
	return u.userRepo.List()
}
