package service

import (
	"github.com/NicholasLiem/ModulAjar_Backend/internal/datastruct"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/dto"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(user datastruct.UserModel) error
	UpdateUser(user dto.UpdateUserDTO) (*datastruct.UserModel, error)
	DeleteUser(userID uint) (*datastruct.UserModel, error)
	GetUser(userID uint) (*datastruct.UserModel, error)
}

type userService struct {
	dao repository.DAO
}

func NewUserService(dao repository.DAO) UserService {
	return &userService{dao: dao}
}

func (u *userService) CreateUser(user datastruct.UserModel) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	_, err = u.dao.NewUserQuery().CreateUser(user)
	return err
}

func (u *userService) UpdateUser(user dto.UpdateUserDTO) (*datastruct.UserModel, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)
	updatedUser, err := u.dao.NewUserQuery().UpdateUser(user)
	return updatedUser, err
}

func (u *userService) DeleteUser(userID uint) (*datastruct.UserModel, error) {
	userData, err := u.dao.NewUserQuery().DeleteUser(userID)
	return userData, err
}

func (u *userService) GetUser(userID uint) (*datastruct.UserModel, error) {
	userData, err := u.dao.NewUserQuery().GetUser(userID)
	return userData, err
}
