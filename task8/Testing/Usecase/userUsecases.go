package Usecase

import (
	"errors"
	"fmt"

	"task1.go/task8/Testing/Domain"
)

type UserUsecase struct {
	UserRepo Domain.UserRepository
}

func NewUserUsecase(userRepo Domain.UserRepository) *UserUsecase {
	return &UserUsecase{UserRepo: userRepo}
}

func (uc *UserUsecase) RegisterUser(user Domain.User) error {
	if user.Username == "" || user.Password == "" {
		return errors.New("missing required fields")
	}

	_, err := uc.UserRepo.AddUser(user)
	return err
}

func (uc *UserUsecase) LoginUser(username string, password string) (Domain.User, error) {
	if username == "" || password == "" {
		return Domain.User{}, errors.New("missing required fields")
	}
	fmt.Println(username, password, "From user usecase")
	user, err := uc.UserRepo.LoginUser(username, password)
	return user, err
}

func (uc *UserUsecase) GetAllUsers() ([]Domain.User, error) {
	users, err := uc.UserRepo.GetAllUsers()
	return users, err
}

func (uc *UserUsecase) DeleteUserID(id string) (Domain.User, error) {
	user, err := uc.UserRepo.DeleteUserID(id)
	return user, err
}

func (uc *UserUsecase) GetUserByID(id string) (Domain.User, error) {
	user, err := uc.UserRepo.GetUserByID(id)

	return user, err
}
