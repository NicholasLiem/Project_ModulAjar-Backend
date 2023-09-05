package repository

import (
	"github.com/NicholasLiem/ModulAjar_Backend/internal/datastruct"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DocumentQuery interface {
	CreateNewDocument(documentPath string) error
	GetUserDocuments(userReferId uint) (*[]datastruct.Document, error)
	GetDocumentById(documentId string) (*datastruct.Document, error)
	DeleteDocument(documentId string) error
}

type documentQuery struct {
	pgdb *gorm.DB
}

func NewDocumentQuery(db *gorm.DB) DocumentQuery {
	return &documentQuery{
		pgdb: db,
	}
}

func (d *documentQuery) CreateNewDocument(documentPath string) error {
	newDocumentId := uuid.New().String()

	newDocument := datastruct.Document{
		DocumentID:   newDocumentId,
		DocumentPath: documentPath,
	}
	if err := d.pgdb.Create(&newDocument).Error; err != nil {
		return err
	}

	return nil

}

func (d *documentQuery) GetUserDocuments(userReferId uint) (*[]datastruct.Document, error) {
	var targetDocuments []datastruct.Document
	err := d.pgdb.Model(datastruct.Document{}).Where("user_refer = ?", userReferId).Find(&targetDocuments).Error
	if err != nil {
		return nil, err
	}

	return &targetDocuments, nil
}

func (d *documentQuery) GetDocumentById(documentId string) (*datastruct.Document, error) {
	var targetDocument datastruct.Document
	err := d.pgdb.Model(datastruct.Document{}).Where("document_id = ?", documentId).First(&targetDocument).Error
	if err != nil {
		return nil, err
	}

	return &targetDocument, nil
}

func (d *documentQuery) DeleteDocument(documentId string) error {
	var targetDocument datastruct.Document
	err := d.pgdb.Model(datastruct.Document{}).Where("document_id = ?", documentId).First(&targetDocument).Error
	if err != nil {
		return err
	}

	err = d.pgdb.Unscoped().Where("document_id", documentId).Delete(&targetDocument).Error
	if err != nil {
		return err
	}

	return nil
}
