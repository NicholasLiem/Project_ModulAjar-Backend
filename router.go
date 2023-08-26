package main

import (
	"github.com/NicholasLiem/ModulAjar_Backend/auth"
	"github.com/NicholasLiem/ModulAjar_Backend/middleware"
	"github.com/NicholasLiem/ModulAjar_Backend/pdfgenerator"
	customRouter "github.com/NicholasLiem/ModulAjar_Backend/router"
	"github.com/NicholasLiem/ModulAjar_Backend/user"
	"github.com/NicholasLiem/ModulAjar_Backend/utils/jwt"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter()

	customRouter.AppRoutes = append(customRouter.AppRoutes, user.Routes, jwt.Routes, pdfgenerator.Routes, auth.Routes)
	for _, route := range customRouter.AppRoutes {

		//create sub route
		routePrefix := router.PathPrefix(route.Prefix).Subrouter()

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

	return router
}
