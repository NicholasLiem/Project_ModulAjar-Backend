package routers

import "github.com/NicholasLiem/ModulAjar_Backend/infrastucture/router"

func AuthRoutes() router.RoutePrefix {
	return router.RoutePrefix{
		Prefix: "/v1/auth",
		SubRoutes: []router.Route{
			{
				"Login",
				"POST",
				"/login",
				nil,
				false,
			},
			{
				"Register",
				"POST",
				"/register/{user_id}",
				nil,
				false,
			},
		},
	}
}
