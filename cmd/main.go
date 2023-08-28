package main

import (
	"github.com/NicholasLiem/ModulAjar_Backend/database"
	"github.com/NicholasLiem/ModulAjar_Backend/infrastucture"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/datastruct"
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

	router := infrastucture.NewRouter()

	log.Println("Running the server on port " + port)

	log.Fatal(http.ListenAndServe(":"+port, router))
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
