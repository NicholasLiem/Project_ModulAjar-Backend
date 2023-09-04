package app

import "github.com/NicholasLiem/ModulAjar_Backend/internal/service"

type MicroserviceServer struct {
	userService            service.UserService
	authService            service.AuthService
	sessionService         service.SessionService
	inputSuggestionService service.InputSuggestionService
}

func NewMicroservice(
	userService service.UserService,
	authService service.AuthService,
	sessionService service.SessionService,
	inputSuggestionService service.InputSuggestionService) *MicroserviceServer {
	return &MicroserviceServer{
		userService:            userService,
		authService:            authService,
		sessionService:         sessionService,
		inputSuggestionService: inputSuggestionService,
	}
}
