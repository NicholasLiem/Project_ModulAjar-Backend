package repository

import (
	"context"
	"encoding/json"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/datastruct"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"strconv"
	"time"
)

type SessionManager interface {
	CreateUserSession(user datastruct.UserModel) error
	GetUserSession(userId string) (bool, error)
}

type sessionManager struct {
	redisClient *redis.Client
}

func NewSessionManager(client *redis.Client) SessionManager {
	return &sessionManager{
		redisClient: client,
	}
}

func (s *sessionManager) CreateUserSession(user datastruct.UserModel) error {
	sessionUser := datastruct.SessionUserClient{
		Authenticated: true,
	}

	sessionUserJSON, err := json.Marshal(sessionUser)
	if err != nil {
		return err
	}

	sessionTimeStr := os.Getenv("SESSION_TIME")
	sessionTime, err := strconv.ParseUint(sessionTimeStr, 10, 64)
	if err != nil {
		return err
	}

	err = s.redisClient.Set(context.Background(), strconv.Itoa(int(user.UserID)), sessionUserJSON, time.Duration(sessionTime)*time.Second).Err()

	return nil
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
