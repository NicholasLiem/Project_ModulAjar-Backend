package routers

import (
	"github.com/NicholasLiem/ModulAjar_Backend/internal/app"
)

func DocumentRoute(server app.MicroserviceServer) RoutePrefix {
	return RoutePrefix{
		Prefix: "/v1/pdf",
		SubRoutes: []Route{
			{
				"Generate PDF",
				"POST",
				"/generate",
				server.Logout,
				true,
			},
		},
	}
}
