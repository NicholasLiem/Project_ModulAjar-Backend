package app

import "github.com/NicholasLiem/ModulAjar_Backend/internal/service"

type MicroserviceServer struct {
	userService service.UserService
}

func NewMicroservice(
	userService service.UserService) *MicroserviceServer {
	return &MicroserviceServer{
		userService: userService,
	}
}
