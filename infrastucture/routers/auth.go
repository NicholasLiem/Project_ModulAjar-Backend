package routers

import (
	"github.com/NicholasLiem/ModulAjar_Backend/infrastucture/router"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/service"
)

var AuthRoutes = router.RoutePrefix{
	Prefix: "/v1/auth",
	SubRoutes: []router.Route{
		{
			"Login",
			"POST",
			"/login",
			service.LoginHandler,
			false,
		},
		{
			"Register",
			"POST",
			"/register/{user_id}",
			service.RegisterHandler,
			false,
		},
	},
}
