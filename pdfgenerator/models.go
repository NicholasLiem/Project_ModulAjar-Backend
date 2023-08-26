package pdfgenerator

import (
	"github.com/NicholasLiem/GoLang_Microservice/database"
	"gorm.io/gorm"
)

type Document struct {
	gorm.Model
	DocumentID   uint    `gorm:"primary_key" json:"document_id,omitempty"`
	UserRefer    uint    `json:"user_refer,omitempty"`
	DocumentPath *string `gorm:"column:document_path" json:"document_path,omitempty"`
}

func FindOneDocumeent(condition interface{}) (Document, error) {
	db := database.DB
	var model Document
	err := db.Where(condition).First(&model).Error
	return model, err
}
