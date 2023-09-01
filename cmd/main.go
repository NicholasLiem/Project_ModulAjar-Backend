package main

import (
	"github.com/NicholasLiem/ModulAjar_Backend/infrastucture"
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
	Loading .env file
	*/
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	/**
	Setting up DB
	*/
	repository.DB = repository.SetupDB()

	/**
	Registering DAO's and Services
	*/
	dao := repository.NewDAO()
	userService := service.NewUserService(dao)
	authService := service.NewAuthService(dao)

	/**
	Registering Services to Server
	*/
	server := app.NewMicroservice(
		userService,
		authService,
	)

	/**
	Run DB Migration
	*/
	Migrate(repository.DB)

	/**
	Setting up the router
	*/
	serverRouter := infrastucture.NewRouter(*server)

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
