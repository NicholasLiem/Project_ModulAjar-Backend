package user

import "github.com/NicholasLiem/GoLang_Microservice/router"

var Routes = router.RoutePrefix{
	Prefix: "/v1/user",
	SubRoutes: []router.Route{
		{
			"Hello",
			"GET",
			"",
			HelloHandler,
			true,
		},
		{
			"Create a new user",
			"POST",
			"/{user_id}",
			CreateUserHandler,
			false,
		},
	},
}
