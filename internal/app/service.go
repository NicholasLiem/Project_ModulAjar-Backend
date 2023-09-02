package app

import "github.com/NicholasLiem/ModulAjar_Backend/internal/service"

type MicroserviceServer struct {
	userService    service.UserService
	authService    service.AuthService
	sessionService service.SessionService
}

func NewMicroservice(
	userService service.UserService,
	authService service.AuthService,
	sessionService service.SessionService) *MicroserviceServer {
	return &MicroserviceServer{
		userService:    userService,
		authService:    authService,
		sessionService: sessionService,
	}
}
