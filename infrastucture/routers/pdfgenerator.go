package routers

import "github.com/NicholasLiem/ModulAjar_Backend/infrastucture/router"

func DocumentRoute() router.RoutePrefix {
	return router.RoutePrefix{
		Prefix: "/v1/pdf",
		SubRoutes: []router.Route{
			{
				"Generate PDF",
				"POST",
				"/generate",
				nil,
				false,
			},
		},
	}
}
