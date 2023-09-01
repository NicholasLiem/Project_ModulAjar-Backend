package service

import (
	"errors"
	"fmt"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/datastruct"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/dto"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/repository"
	"github.com/NicholasLiem/ModulAjar_Backend/utils"
	jwt2 "github.com/NicholasLiem/ModulAjar_Backend/utils/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	SignIn(loginDTO dto.LoginDTO) (*string, error)
	SignUp(model datastruct.UserModel) (*string, error)
	//LogOut(userID uint) error
}

type authService struct {
	dao repository.DAO
}

func NewAuthService(dao repository.DAO) AuthService {
	return &authService{dao: dao}
}

func (a *authService) SignIn(loginDTO dto.LoginDTO) (*string, error) {
	password, err := a.dao.NewUserQuery().GetUserPasswordByEmail(loginDTO.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(*password), []byte(loginDTO.Password))
	if err != nil {
		return nil, fmt.Errorf("passwords dont match %v", err)
	} else {
		jwt, err := jwt2.CreateJWT()
		if err != nil {
			return nil, err
		}

		return &jwt.Token, nil
	}
}

func (a *authService) SignUp(model datastruct.UserModel) (*string, error) {

	if !utils.IsEmailValid(model.Email) {
		return nil, errors.New("email is not valid")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(model.Password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	model.Password = string(hashedPassword)

	_, err = a.dao.NewUserQuery().CreateUser(model)
	if err != nil {
		return nil, err
	}

	jwt, err := jwt2.CreateJWT()
	if err != nil {
		return nil, err
	}

	return &jwt.Token, nil
}

//func hasValidToken(r *http.Request) bool {
//	tokenStr := middleware.DecodeJWTToken(r)
//	claims, err := jwt2.VerifyJWT(tokenStr)
//
//	if err != nil || claims == nil {
//		return false
//	}
//
//	return true
//}
//
