package pdfgenerator

import "github.com/NicholasLiem/GoLang_Microservice/router"

var Routes = router.RoutePrefix{
	Prefix: "/v1/pdf",
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
