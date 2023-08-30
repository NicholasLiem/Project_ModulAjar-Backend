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

//func FindUserByIdHandler(rw http.ResponseWriter, r *http.Request) {
//	params := mux.Vars(r)
//	id := params["user_id"]
//
//	userID, err := VerifyUserId(id)
//	if err != nil {
//		response.ErrorResponse(rw, http.StatusBadRequest, "Invalid user ID")
//		return
//	}
//
//	condition := datastruct.UserModel{UserID: uint(userID)}
//
//	foundUser, err := datastruct.FindOneUser(condition)
//	if err != nil {
//		response.ErrorResponse(rw, http.StatusNotFound, "User not found")
//		return
//	}
//
//	response.SuccessResponse(rw, http.StatusOK, "User found", foundUser)
//	return
//}

//func DeleteUserByIdHandler(rw http.ResponseWriter, r *http.Request) {
//	params := mux.Vars(r)
//	id := params["user_id"]
//
//	userID, err := VerifyUserId(id)
//	if err != nil {
//		response.ErrorResponse(rw, http.StatusBadRequest, "Invalid user ID")
//		return
//	}
//
//	condition := datastruct.UserModel{UserID: uint(userID)}
//
//	foundUser, err := datastruct.FindOneUser(condition)
//	if err != nil {
//		response.ErrorResponse(rw, http.StatusNotFound, "User not found")
//		return
//	}
//
//	userData := foundUser
//
//	err = foundUser.Delete()
//	if err != nil {
//		response.ErrorResponse(rw, http.StatusInternalServerError, "Fail to delete user")
//		return
//	}
//
//	response.SuccessResponse(rw, http.StatusOK, "User deleted", userData)
//	return
//}
//
