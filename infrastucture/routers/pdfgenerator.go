package routers

import (
	"github.com/NicholasLiem/ModulAjar_Backend/infrastucture/router"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/service"
)

var DocumentRoutes = router.RoutePrefix{
	Prefix: "/v1/pdf",
	SubRoutes: []router.Route{
		{
			"Generate PDF Handler",
			"POST",
			"/generate",
			service.GenPDFHandler,
			false,
		},
	},
}
