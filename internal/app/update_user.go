package app

import (
	"encoding/json"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/dto"
	response "github.com/NicholasLiem/ModulAjar_Backend/utils/http"
	"net/http"
)

func (m *MicroserviceServer) UpdateUser(w http.ResponseWriter, r *http.Request) {

	var updateUser dto.UpdateUserDTO
	err := json.NewDecoder(r.Body).Decode(&updateUser)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "Fail to update user "+err.Error())
		return
	}

	updatedUser, err := m.userService.UpdateUser(updateUser)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, "Fail to update user "+err.Error())
		return
	}

	response.SuccessResponse(w, http.StatusOK, "Successfully updated the user", updatedUser)
	return
}
