package routers

import (
	"github.com/NicholasLiem/ModulAjar_Backend/infrastucture/router"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/app"
)

func AuthRoutes(server app.MicroserviceServer) router.RoutePrefix {
	return router.RoutePrefix{
		Prefix: "/v1/auth",
		SubRoutes: []router.Route{
			{
				"Login",
				"POST",
				"/login",
				server.Login,
				false,
			},
			{
				"Register",
				"POST",
				"/register",
				server.Register,
				false,
			},
			{
				"Logout",
				"POST",
				"/logout",
				server.Logout,
				true,
			},
		},
	}
}
