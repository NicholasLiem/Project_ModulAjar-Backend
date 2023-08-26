package user

import "github.com/NicholasLiem/ModulAjar_Backend/router"

var Routes = router.RoutePrefix{
	Prefix: "/v1/user",
	SubRoutes: []router.Route{
		{
			"Create a new user",
			"POST",
			"/{user_id}",
			CreateUserHandler,
			false,
		},
		{
			"Find a user by id",
			"GET",
			"/{user_id}",
			FindUserByIdHandler,
			true,
		},
		{
			"Delete user by id",
			"DELETE",
			"/{user_id}",
			DeleteUserByIdHandler,
			true,
		},
		{
			"Update user by id",
			"PUT",
			"/{user_id}",
			UpdateUserHandler,
			true,
		},
	},
}
