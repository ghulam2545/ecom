package service

import (
	"ecom/model"
	"ecom/repo"
	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
	"time"
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

func (u *UserService) Signup(request *model.SignupRequest) (*model.User, error) {
	email := request.Email
	password := request.Password
	role := request.Role

	h, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := model.User{UserId: xid.New().String(), Email: email, Password: string(h), Role: role, CreatedAt: time.Now()}
	inserted, err := u.userRepo.Save(&user)
	if err != nil {
		return nil, err
	}
	return inserted, nil
}

func (u *UserService) Login(request *model.LoginRequest) (string, *model.User, error) {
	email := request.Email
	password := request.Password

	user, err := u.userRepo.GetByEmail(email)
	if err != nil {
		return "", nil, err
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return "", nil, err
	}

	token, err := GenerateToken(user.UserId, user.Email, user.Role, 24*time.Hour)
	return token, user, err
}
