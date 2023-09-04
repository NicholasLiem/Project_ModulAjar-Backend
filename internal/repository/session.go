package repository

import (
	"context"
	"encoding/json"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/datastruct"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"strconv"
	"time"
)

type SessionManager interface {
	CreateUserSession(user datastruct.UserModel) (*string, error)
	GetUserSession(userId string) (bool, error)
	DeleteUserSession(userId string) error
	GetSessionData(sessionId string) (*datastruct.SessionUserClient, error)
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
	sessionId := uuid.New().String()

	sessionUser := datastruct.SessionUserClient{
		UserID:        strconv.Itoa(int(user.UserID)),
		Role:          user.Role,
		Authenticated: true,
	}

	sessionUserJSON, err := json.Marshal(sessionUser)
	if err != nil {
		return nil, err
	}

	sessionTimeStr := os.Getenv("SESSION_TIME")
	sessionTime, err := strconv.ParseUint(sessionTimeStr, 10, 64)
	if err != nil {
		return nil, err
	}

	err = s.redisClient.Set(context.Background(), sessionId, sessionUserJSON, time.Duration(sessionTime)*time.Second).Err()

	return &sessionId, nil
}

func (s *sessionManager) GetUserSession(userId string) (bool, error) {
	sessionUserJSON, err := s.redisClient.Get(context.Background(), userId).Result()
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
func (s *sessionManager) GetSessionData(sessionId string) (*datastruct.SessionUserClient, error) {
	sessionUserJSON, err := s.redisClient.Get(context.Background(), sessionId).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, err
		}
		log.Printf("Failed to get session data from redis: %s", err.Error())
		return nil, err
	}

	var sessionData datastruct.SessionUserClient
	if err := json.Unmarshal([]byte(sessionUserJSON), &sessionData); err != nil {
		log.Printf("Failed to unmarshal session data: %s", err.Error())
		return nil, err
	}

	return &sessionData, nil
}

func (s *sessionManager) DeleteUserSession(userId string) error {
	_, err := s.redisClient.Del(context.Background(), userId).Result()
	if err != nil {
		log.Fatalf("Failed to delete key: %v", err)
	}
	return nil
}
