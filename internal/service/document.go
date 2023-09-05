package service

import "github.com/NicholasLiem/ModulAjar_Backend/internal/repository"

type DocumentService interface {
	GenerateNewDocument(documentDTO string) (*string, error)
}

type documentService struct {
	dao repository.DAO
}

func NewDocumentService(dao repository.DAO) DocumentService {
	return &documentService{
		dao: dao,
	}
}

func (d *documentService) GenerateNewDocument(documentDTO string) (*string, error) {
	return nil, nil
}
