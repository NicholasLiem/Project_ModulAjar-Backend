package auth

import "github.com/NicholasLiem/GoLang_Microservice/router"

var Routes = router.RoutePrefix{
	Prefix: "/v1/auth",
	SubRoutes: []router.Route{
		{
			"Login",
			"POST",
			"/login",
			nil,
			false,
		},
	},
}
