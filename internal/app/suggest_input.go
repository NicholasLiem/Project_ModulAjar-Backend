package app

import (
	"encoding/json"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/dto"
	response "github.com/NicholasLiem/ModulAjar_Backend/utils/http"
	"github.com/NicholasLiem/ModulAjar_Backend/utils/messages"
	"net/http"
)

func (m *MicroserviceServer) InputSuggestion(w http.ResponseWriter, r *http.Request) {

	var newSuggestion dto.InputSuggestionDTO
	err := json.NewDecoder(r.Body).Decode(&newSuggestion)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, messages.InvalidRequestData)
		return
	}

	responseMsg, err := m.inputSuggestionService.SuggestIdeas(newSuggestion)
	if err != nil {
		return
	}

	response.SuccessResponse(w, http.StatusOK, "Suggestion achieved", responseMsg)
	return
}
