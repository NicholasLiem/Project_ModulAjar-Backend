package jwt

import (
	"github.com/NicholasLiem/ModulAjar_Backend/infrastucture/router"
)

var Routes = router.RoutePrefix{
	Prefix: "/jwt",
	SubRoutes: []router.Route{
		{
			"GetJWT",
			"GET",
			"",
			GetJwt,
			false,
		},
	},
}
