package auth

import "github.com/NicholasLiem/ModulAjar_Backend/router"

var Routes = router.RoutePrefix{
	Prefix: "/v1/auth",
	SubRoutes: []router.Route{
		{
			"Login",
			"POST",
			"/login",
			LoginHandler,
			false,
		},
		{
			"Register",
			"POST",
			"/register/{user_id}",
			RegisterHandler,
			false,
		},
	},
}
