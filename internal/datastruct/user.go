package datastruct

import (
	"errors"
	"github.com/NicholasLiem/ModulAjar_Backend/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	UserID       uint       `gorm:"primaryIndex:user_id" json:"user_id,omitempty"`
	Username     string     `gorm:"column:username" json:"username,omitempty"`
	Email        string     `gorm:"column:email;unique_index" json:"email,omitempty"`
	PasswordHash string     `gorm:"column:password;not null" json:"password_hash,omitempty"`
	Documents    []Document `gorm:"foreignKey:UserRefer;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"documents,omitempty"`
}

func (model *UserModel) SetPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password can not be empty")
	}

	bytePassword := []byte(password)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	model.PasswordHash = string(passwordHash)
	return nil
}

func (model *UserModel) CheckPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(model.PasswordHash)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

func FindOneUser(condition interface{}) (UserModel, error) {
	db := database.DB
	var model UserModel
	err := db.Where(condition).First(&model).Error
	return model, err
}

func SaveOneUser(data interface{}) error {
	db := database.DB
	err := db.Save(data).Error
	return err
}

func CreateUser(data interface{}) error {
	db := database.DB
	err := db.Create(data).Error
	return err
}

func (model *UserModel) Update(data interface{}) error {
	db := database.DB
	err := db.Model(model).Updates(data).Error
	return err
}

func (model *UserModel) Delete() error {
	db := database.DB
	return db.Delete(model).Error
}

func (model *UserModel) AddDocument(data interface{}) error {
	db := database.DB

	document := Document{
		UserRefer: model.UserID,
	}

	dataMap, ok := data.(map[string]interface{})
	if !ok {
		return errors.New("invalid data format")
	}

	if documentPath, ok := dataMap["DocumentPath"]; ok {
		if pathStr, ok := documentPath.(string); ok {
			document.DocumentPath = &pathStr
		}
	}

	if err := db.Create(&document).Error; err != nil {
		return errors.New("fail to create the document")
	}

	return nil
}

func (model *UserModel) GetDocuments() []Document {
	db := database.DB

	var documents []Document
	err := db.Model(model).Association("Documents").Find(&documents)

	if err != nil {
		return nil
	}

	return documents
}
