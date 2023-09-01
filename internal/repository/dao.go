package repository

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

type DAO interface {
	NewUserQuery() UserQuery
}

type dao struct {
	db    *gorm.DB
	redis *redis.Client
}

var DisableLogger = logger.New(
	nil, // Use the default logger output, which is discarded
	logger.Config{
		LogLevel: logger.Silent,
	},
)

func NewDAO(db *gorm.DB, rc *redis.Client) DAO {
	return &dao{
		db:    db,
		redis: rc,
	}
}

func SetupDB() *gorm.DB {
	var dbHost = os.Getenv("DB_HOST")
	var dbName = os.Getenv("DB_NAME")
	var dbUsername = os.Getenv("DB_USERNAME")
	var dbPassword = os.Getenv("DB_PASSWORD")
	var dbPort = os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUsername, dbPassword, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: DisableLogger,
	})

	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("Failed to get DB instance: " + err.Error())
	}

	log.Println("Successfully connected to DB")

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return db
}

func SetupRedis(ctx context.Context) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       10,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic("Failed to connect to Redis: " + err.Error())
	}

	log.Println("Successfully connected to Redis")

	return client
}

func (d *dao) NewUserQuery() UserQuery {
	return NewUserQuery(d.db)
}
