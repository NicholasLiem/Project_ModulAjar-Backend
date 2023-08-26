package main

import (
	"github.com/NicholasLiem/GoLang_Microservice/database"
	"github.com/NicholasLiem/GoLang_Microservice/pdfgenerator"
	"github.com/NicholasLiem/GoLang_Microservice/user"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.DB = database.SetupDB()

	Migrate(database.DB)

	port := os.Getenv("PORT")

	router := NewRouter()

	log.Println("Running the server on port " + port)

	log.Fatal(http.ListenAndServe("127.0.0.1:"+port, router))
}

func Migrate(db *gorm.DB) {
	errPdf := db.AutoMigrate(&pdfgenerator.Document{})
	if errPdf != nil {
		return
	}
	err := db.AutoMigrate(&user.UserModel{})
	if err != nil {
		return
	}
}
