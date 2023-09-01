package repository

import (
	"github.com/NicholasLiem/ModulAjar_Backend/internal/datastruct"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/dto"
)

type UserQuery interface {
	CreateUser(user datastruct.UserModel) (*uint, error)
	UpdateUser(user dto.UpdateUserDTO) (*datastruct.UserModel, error)
	DeleteUser(userID uint) (*datastruct.UserModel, error)
	GetUser(userID uint) (*datastruct.UserModel, error)
}

type userQuery struct{}

func (u *userQuery) CreateUser(user datastruct.UserModel) (*uint, error) {
	db := DB

	newUser := datastruct.UserModel{
		UserID:   user.UserID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

	if err := db.Create(&newUser).Error; err != nil {
		return nil, err
	}

	return &newUser.ID, nil
}

func (u *userQuery) UpdateUser(user dto.UpdateUserDTO) (*datastruct.UserModel, error) {
	db := DB
	err := db.Model(datastruct.UserModel{}).Where("user_id = ?", user.UserID).Updates(user).Error

	var updatedUser datastruct.UserModel
	err = db.Where("user_id = ?", user.UserID).First(&updatedUser).Error
	if err != nil {
		return nil, err
	}

	return &updatedUser, err
}

func (u *userQuery) DeleteUser(userID uint) (*datastruct.UserModel, error) {
	db := DB
	var userData datastruct.UserModel
	err := db.Model(datastruct.UserModel{}).Where("user_id = ?", userID).First(&userData).Error
	if err != nil {
		return nil, err
	}

	/**
	Perform hard delete, if you want to soft delete, delete the Unscoped function
	*/
	err = db.Unscoped().Where("user_id = ?", userID).Delete(&userData).Error
	if err != nil {
		return nil, err
	}

	return &userData, err
}

func (u *userQuery) GetUser(userID uint) (*datastruct.UserModel, error) {
	db := DB

	var userData datastruct.UserModel
	err := db.Where("user_id = ?", userID).First(&userData).Error
	return &userData, err
}

//func (model *UserModel) AddDocument(data interface{}) error {
//	db := database.DB
//
//	document := Document{
//		UserRefer: model.UserID,
//	}
//
//	dataMap, ok := data.(map[string]interface{})
//	if !ok {
//		return errors.New("invalid data format")
//	}
//
//	if documentPath, ok := dataMap["DocumentPath"]; ok {
//		if pathStr, ok := documentPath.(string); ok {
//			document.DocumentPath = &pathStr
//		}
//	}
//
//	if err := db.Create(&document).Error; err != nil {
//		return errors.New("fail to create the document")
//	}
//
//	return nil
//}
//
//func (model *UserModel) GetDocuments() []Document {
//	db := database.DB
//
//	var documents []Document
//	err := db.Model(model).Association("Documents").Find(&documents)
//
//	if err != nil {
//		return nil
//	}
//
//	return documents
//}
