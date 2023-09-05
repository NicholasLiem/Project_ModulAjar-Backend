package routers

import (
	"github.com/NicholasLiem/ModulAjar_Backend/internal/app"
)

func InputSuggestionRoutes(server app.MicroserviceServer) RoutePrefix {
	return RoutePrefix{
		Prefix: "/v1/suggest",
		SubRoutes: []Route{
			{
				"Create a new prompt",
				"POST",
				"/",
				server.InputSuggestion,
				true,
			},
		},
	}
}
