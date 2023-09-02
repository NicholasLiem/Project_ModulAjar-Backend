package service

import (
	"errors"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/datastruct"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/repository"
	"strconv"
)

type SessionService interface {
	CreateUserSession(user datastruct.UserModel) error
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

func (s *sessionService) CreateUserSession(user datastruct.UserModel) error {
	isLoggedIn, err := s.dao.NewSessionManager().GetUserSession(strconv.Itoa(int(user.UserID)))
	if err != nil {
		return err
	}

	if isLoggedIn == true {
		return errors.New("already logged in")
	}

	return s.dao.NewSessionManager().CreateUserSession(user)
}

func (s *sessionService) GetUserSession(userId string) (bool, error) {
	return s.dao.NewSessionManager().GetUserSession(userId)
}
