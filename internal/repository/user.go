package repository

import (
	"context"
	"encoding/json"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/datastruct"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/dto"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"log"
)

type UserQuery interface {
	CreateUser(user datastruct.UserModel) (*datastruct.UserModel, error)
	UpdateUser(user dto.UpdateUserDTO) (*datastruct.UserModel, error)
	DeleteUser(userID uint) (*datastruct.UserModel, error)
	GetUser(userID uint) (*datastruct.UserModel, error)
	GetUserPasswordByEmail(email string) (*string, error)
	GetUserByEmail(email string) (*datastruct.UserModel, error)
}

type userQuery struct {
	pgdb  *gorm.DB
	redis *redis.Client
}

func NewUserQuery(pgdb *gorm.DB, redis *redis.Client) UserQuery {
	return &userQuery{
		pgdb:  pgdb,
		redis: redis,
	}
}

func (u *userQuery) CreateUser(user datastruct.UserModel) (*datastruct.UserModel, error) {
	newUser := datastruct.UserModel{
		UserID:    user.UserID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	}

	if err := u.pgdb.Create(&newUser).Error; err != nil {
		return nil, err
	}
	return &newUser, nil
}

func (u *userQuery) UpdateUser(user dto.UpdateUserDTO) (*datastruct.UserModel, error) {
	err := u.pgdb.Model(datastruct.UserModel{}).Where("user_id = ?", user.UserID).Updates(user).Error

	var updatedUser datastruct.UserModel
	err = u.pgdb.Where("user_id = ?", user.UserID).First(&updatedUser).Error
	if err != nil {
		return nil, err
	}

	return &updatedUser, err
}

func (u *userQuery) DeleteUser(userID uint) (*datastruct.UserModel, error) {
	var userData datastruct.UserModel
	err := u.pgdb.Model(datastruct.UserModel{}).Where("user_id = ?", userID).First(&userData).Error
	if err != nil {
		return nil, err
	}

	/**
	Perform hard delete, if you want to soft delete, delete the Unscoped function
	*/
	err = u.pgdb.Unscoped().Where("user_id = ?", userID).Delete(&userData).Error
	if err != nil {
		return nil, err
	}

	return &userData, err
}

func (u *userQuery) GetUser(userID uint) (*datastruct.UserModel, error) {
	var userData datastruct.UserModel
	err := u.pgdb.Where("user_id = ?", userID).First(&userData).Error
	return &userData, err
}

func (u *userQuery) GetUserPasswordByEmail(email string) (*string, error) {
	var password string
	err := u.pgdb.Model(&datastruct.UserModel{}).Where("email = ?", email).Select("password").Scan(&password).Error
	return &password, err
}

func (u *userQuery) GetUserByEmail(email string) (*datastruct.UserModel, error) {
	userDataFromRedis, err := u.redis.Get(context.Background(), email).Result()
	if err == nil {
		var userData datastruct.UserModel
		if err := json.Unmarshal([]byte(userDataFromRedis), &userData); err != nil {
			return nil, err
		}
		return &userData, nil
	}

	var userData datastruct.UserModel
	err = u.pgdb.Where("email = ?", email).First(&userData).Error
	if err != nil {
		return nil, err
	}

	userDataJSON, err := json.Marshal(userData)
	if err != nil {
		return nil, err
	}

	err = u.redis.Set(context.Background(), email, userDataJSON, 0).Err()
	if err != nil {
		log.Printf("Failed to set user data in Redis: %s", err.Error())
	}

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
