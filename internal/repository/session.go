package repository

import (
	"context"
	"encoding/json"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/datastruct"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type SessionManager interface {
	CreateUserSession(user datastruct.UserModel) (*string, error)
	GetUserSession(sessionID string) (bool, error)
}

type sessionManager struct {
	redisClient *redis.Client
}

func NewSessionManager(client *redis.Client) SessionManager {
	return &sessionManager{
		redisClient: client,
	}
}

func (s *sessionManager) CreateUserSession(user datastruct.UserModel) (*string, error) {
	sessionID := uuid.New().String()

	sessionUser := datastruct.SessionUserClient{
		UserID:        user.UserID,
		Authenticated: true,
		Role:          string(user.Role),
	}

	sessionUserJSON, err := json.Marshal(sessionUser)
	if err != nil {
		return nil, err
	}

	err = s.redisClient.Set(context.Background(), sessionID, sessionUserJSON, 180*time.Second).Err()
	if err != nil {
		log.Printf("Failed to set session data in redis %s", err.Error())
	}

	return &sessionID, nil
}

func (s *sessionManager) GetUserSession(sessionID string) (bool, error) {
	sessionUserJSON, err := s.redisClient.Get(context.Background(), sessionID).Result()
	if err != nil {
		if err == redis.Nil {
			return false, nil
		}
		log.Printf("Failed to get session data from redis: %s", err.Error())
		return false, err
	}

	var sessionUser datastruct.SessionUserClient
	if err := json.Unmarshal([]byte(sessionUserJSON), &sessionUser); err != nil {
		log.Printf("Failed to unmarshal session data: %s", err.Error())
		return false, err
	}

	return sessionUser.Authenticated, nil
}
