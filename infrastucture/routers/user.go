package routers

import (
	"github.com/NicholasLiem/ModulAjar_Backend/infrastucture/router"
)

var UserRoutes = router.RoutePrefix{
	Prefix: "/v1/user",
	SubRoutes: []router.Route{
		{
			"Create a new user",
			"POST",
			"/{user_id}",
			nil,
			false,
		},
		//{
		//	"Find a user by id",
		//	"GET",
		//	"/{user_id}",
		//	service.FindUserByIdHandler,
		//	true,
		//},
		//{
		//	"Delete user by id",
		//	"DELETE",
		//	"/{user_id}",
		//	service.DeleteUserByIdHandler,
		//	true,
		//},
		//{
		//	"Update user by id",
		//	"PUT",
		//	"/{user_id}",
		//	service.UpdateUserHandler,
		//	true,
		//},
	},
}
