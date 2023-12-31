package infrastucture

import (
	"github.com/NicholasLiem/ModulAjar_Backend/infrastucture/middleware"
	customRouter "github.com/NicholasLiem/ModulAjar_Backend/infrastucture/router"
	"github.com/NicholasLiem/ModulAjar_Backend/infrastucture/routers"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/app"
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
	"net/http"
)

func NewRouter(server app.MicroserviceServer, redisClient redis.Client) *mux.Router {

	router := mux.NewRouter()

	customRouter.AppRoutes = append(customRouter.AppRoutes,
		routers.UserRoutes(server),
		routers.AuthRoutes(server),
		routers.InputSuggestionRoutes(server),
	)

	for _, route := range customRouter.AppRoutes {

		//create sub route
		routePrefix := router.PathPrefix(route.Prefix).Subrouter()

		//for each sub route
		for _, subRoute := range route.SubRoutes {

			var handler http.Handler
			handler = subRoute.HandlerFunc

			if subRoute.Protected {
				handler = middleware.Middleware(subRoute.HandlerFunc, &redisClient) // use middleware
			}

			//register the route
			routePrefix.Path(subRoute.Pattern).Handler(handler).Methods(subRoute.Method).Name(subRoute.Name)
		}

	}

	return router
}
