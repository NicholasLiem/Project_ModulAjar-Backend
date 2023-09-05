package main

import (
	"context"
	"github.com/NicholasLiem/ModulAjar_Backend/infrastucture/routers"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/app"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/datastruct"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/repository"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/service"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {

	/**
	Creating context
	*/
	ctx := context.Background()

	/**
	Loading .env file
	*/
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	/**
	Setting up DB and redis
	*/
	db := repository.SetupDB()
	redis := repository.SetupRedis(ctx)
	openai := repository.SetupOpenAI()

	/**
	Registering DAO's and Services
	*/
	dao := repository.NewDAO(db, redis, openai)
	userService := service.NewUserService(dao)
	authService := service.NewAuthService(dao)
	sessionService := service.NewSessionService(dao)
	inputSuggestionService := service.NewInputSuggestionService(dao)
	documentService := service.NewDocumentService(dao)

	/**
	Registering Services to Server
	*/
	server := app.NewMicroservice(
		userService,
		authService,
		sessionService,
		inputSuggestionService,
		documentService,
	)

	/**
	Run DB Migration
	*/
	Migrate(db)

	/**
	Setting up the router
	*/
	serverRouter := routers.NewRouter(*server, *redis)

	/**
	Running the server
	*/
	port := os.Getenv("PORT")
	log.Println("Running the server on port " + port)

	if os.Getenv("ENVIRONMENT") == "DEV" {
		log.Fatal(http.ListenAndServe("127.0.0.1:"+port, serverRouter))
	}
	log.Fatal(http.ListenAndServe(":"+port, serverRouter))
}

func Migrate(db *gorm.DB) {
	errPdf := db.AutoMigrate(&datastruct.Document{})
	if errPdf != nil {
		return
	}
	err := db.AutoMigrate(&datastruct.UserModel{})
	if err != nil {
		return
	}
}
