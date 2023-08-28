package main

import (
	"github.com/NicholasLiem/ModulAjar_Backend/infrastucture/middleware"
	"github.com/NicholasLiem/ModulAjar_Backend/infrastucture/router"
	customRouter "github.com/NicholasLiem/ModulAjar_Backend/infrastucture/router"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/datastruct"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/handlers"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/repository"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/service"
	"github.com/gorilla/mux"
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

	repository.DB = repository.SetupDB()
	dao := repository.NewDAO()
	userService := service.NewUserService(dao)

	userHandler := handlers.NewUserHandler(userService)

	var userRoutes = router.RoutePrefix{
		Prefix: "/v1/user",
		SubRoutes: []router.Route{
			{
				"Create a new user",
				"POST",
				"/{user_id}",
				userHandler.CreateUser,
				false,
			},
			//{
			//	"Find a user by id",
			//	"GET",
			//	"/{user_id}",
			//	service.FindUserByIdHandler,
			//	true,
			//},
			//{
			//	"Delete user by id",
			//	"DELETE",
			//	"/{user_id}",
			//	service.DeleteUserByIdHandler,
			//	true,
			//},
			//{
			//	"Update user by id",
			//	"PUT",
			//	"/{user_id}",
			//	service.UpdateUserHandler,
			//	true,
			//},
		},
	}

	Migrate(repository.DB)

	port := os.Getenv("PORT")

	//router := infrastucture.NewRouter()

	newRouter := mux.NewRouter()

	customRouter.AppRoutes = append(customRouter.AppRoutes, userRoutes)
	for _, route := range customRouter.AppRoutes {

		//create sub route
		routePrefix := newRouter.PathPrefix(route.Prefix).Subrouter()

		//for each sub route
		for _, subRoute := range route.SubRoutes {

			var handler http.Handler
			handler = subRoute.HandlerFunc

			if subRoute.Protected {
				handler = middleware.Middleware(subRoute.HandlerFunc) // use middleware
			}

			//register the route
			routePrefix.Path(subRoute.Pattern).Handler(handler).Methods(subRoute.Method).Name(subRoute.Name)
		}

	}

	log.Println("Running the server on port " + port)

	log.Fatal(http.ListenAndServe(":"+port, newRouter))
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
