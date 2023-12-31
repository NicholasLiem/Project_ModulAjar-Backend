package service

import (
	"errors"
	"fmt"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/datastruct"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/dto"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/repository"
	"github.com/NicholasLiem/ModulAjar_Backend/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	SignIn(loginDTO dto.LoginDTO) (*datastruct.UserModel, error)
	SignUp(model datastruct.UserModel) (*datastruct.UserModel, error)
	//LogOut(userID uint) error
}

type authService struct {
	dao repository.DAO
}

func NewAuthService(dao repository.DAO) AuthService {
	return &authService{dao: dao}
}

func (a *authService) SignIn(loginDTO dto.LoginDTO) (*datastruct.UserModel, error) {
	password, err := a.dao.NewUserQuery().GetUserPasswordByEmail(loginDTO.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(*password), []byte(loginDTO.Password))
	if err != nil {
		return nil, fmt.Errorf("passwords dont match %v", err)
	} else {
		userData, err := a.dao.NewUserQuery().GetUserByEmail(loginDTO.Email)
		if err != nil {
			return nil, err
		}

		return userData, nil
	}
}

func (a *authService) SignUp(model datastruct.UserModel) (*datastruct.UserModel, error) {

	if !utils.IsEmailValid(model.Email) {
		return nil, errors.New("email is not valid")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(model.Password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	model.Password = string(hashedPassword)

	userData, err := a.dao.NewUserQuery().CreateUser(model)
	if err != nil {
		return nil, err
	}

	return userData, nil
}
