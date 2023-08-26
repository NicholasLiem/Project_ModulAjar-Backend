package pdfgenerator

import "github.com/NicholasLiem/GoLang_Microservice/router"

var Routes = router.RoutePrefix{
	Prefix: "/pdf",
	SubRoutes: []router.Route{
		{
			"Generate PDF Handler",
			"POST",
			"/generate",
			GenPDFHandler,
			false,
		},
	},
}
