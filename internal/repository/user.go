package repository

import (
	"github.com/NicholasLiem/ModulAjar_Backend/internal/datastruct"
)

type UserQuery interface {
	CreateUser(user datastruct.UserModel) (*uint, error)
	UpdateUser(user datastruct.UserModel) error
}

type userQuery struct{}

func (u *userQuery) CreateUser(user datastruct.UserModel) (*uint, error) {
	db := DB

	newUser := datastruct.UserModel{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

	if err := db.Create(&newUser).Error; err != nil {
		return nil, err
	}

	return &newUser.ID, nil
}

func (u *userQuery) UpdateUser(user datastruct.UserModel) error {
	db := DB
	err := db.Model(datastruct.UserModel{}).Where("user_id = ?", user.UserID).Updates(user).Error
	return err
}

//func (u *userQuery) GetUser(condition interface{}) (datastruct.UserModel, error) {
//	db := DB
//	err := db.Where(condition).First().Error
//	return model, err
//}

//func SaveOneUser(data interface{}) error {
//	db := DB
//	err := db.Save(data).Error
//	return err
//}
//
//func CreateUser(data interface{}) error {
//	db := DB
//	err := db.Create(data).Error
//	return err
//}
//
//func (model *UserModel) Update(data interface{}) error {
//	db := DB
//	err := db.Model(model).Updates(data).Error
//	return err
//}
//
//func (model *UserModel) Delete() error {
//	db := DB
//	return db.Delete(model).Error
//}

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
