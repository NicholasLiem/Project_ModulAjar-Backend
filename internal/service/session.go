package service

import (
	"errors"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/datastruct"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/repository"
	"strconv"
)

type SessionService interface {
	CreateUserSession(user datastruct.UserModel) (*string, error)
	GetUserSession(sessionId string) (bool, error)
	InvalidateUserSession(sessionId string) error
	GetSessionData(sessionId string) (*datastruct.SessionUserClient, error)
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
	isLoggedIn, err := s.dao.NewSessionManager().GetUserSession(strconv.Itoa(int(user.UserID)))
	if err != nil {
		return nil, err
	}

	if isLoggedIn == true {
		return nil, errors.New("already logged in")
	}

	return s.dao.NewSessionManager().CreateUserSession(user)
}

func (s *sessionService) GetUserSession(sessionId string) (bool, error) {
	return s.dao.NewSessionManager().GetUserSession(sessionId)
}

func (s *sessionService) GetSessionData(sessionId string) (*datastruct.SessionUserClient, error) {
	return s.dao.NewSessionManager().GetSessionData(sessionId)
}

func (s *sessionService) InvalidateUserSession(sessionId string) error {
	return s.dao.NewSessionManager().DeleteUserSession(sessionId)
}
