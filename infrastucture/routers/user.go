package routers

import (
	"github.com/NicholasLiem/ModulAjar_Backend/infrastucture/router"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/app"
)

func UserRoutes(server app.MicroserviceServer) router.RoutePrefix {
	return router.RoutePrefix{
		Prefix: "/v1/user",
		SubRoutes: []router.Route{
			{
				"Create a new user",
				"POST",
				"/{user_id}",
				server.CreateUser,
				false,
			},
			{
				"Update a user",
				"PUT",
				"/",
				server.UpdateUser,
				false,
			},
		},
	}
}
