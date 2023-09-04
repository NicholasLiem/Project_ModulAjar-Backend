package service

import (
	"github.com/NicholasLiem/ModulAjar_Backend/internal/dto"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/repository"
)

type InputSuggestionService interface {
	SuggestIdeas(prompt dto.InputSuggestionDTO) (*string, error)
}

type inputSuggestionService struct {
	dao repository.DAO
}

func NewInputSuggestionService(dao repository.DAO) InputSuggestionService {
	return &inputSuggestionService{
		dao: dao,
	}
}

func (u *inputSuggestionService) SuggestIdeas(prompt dto.InputSuggestionDTO) (*string, error) {
	resp, err := u.dao.NewInputSuggestionQuery().GeneratePrompt(prompt.Prompt)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
