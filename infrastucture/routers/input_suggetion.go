package routers

import (
	"github.com/NicholasLiem/ModulAjar_Backend/infrastucture/router"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/app"
)

func InputSuggestionRoutes(server app.MicroserviceServer) router.RoutePrefix {
	return router.RoutePrefix{
		Prefix: "/v1/suggest",
		SubRoutes: []router.Route{
			{
				"Create a new prompt",
				"POST",
				"/",
				server.InputSuggestion,
				false,
			},
		},
	}
}
