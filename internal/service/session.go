package service

import (
	"github.com/NicholasLiem/ModulAjar_Backend/internal/datastruct"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/repository"
)

type SessionService interface {
	CreateUserSession(user datastruct.UserModel) (*string, error)
	GetUserSession(sessionID string) (bool, error)
}

type sessionService struct {
	dao repository.DAO
}

func NewSessionService(dao repository.DAO) SessionService {
	return &sessionService{
		dao: dao,
	}
}

func (s *sessionService) CreateUserSession(user datastruct.UserModel) (*string, error) {
	return s.dao.NewSessionManager().CreateUserSession(user)
}

func (s *sessionService) GetUserSession(sessionID string) (bool, error) {
	return s.dao.NewSessionManager().GetUserSession(sessionID)
}
