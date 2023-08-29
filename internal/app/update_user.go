package app

import (
	"encoding/json"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/datastruct"
	"github.com/NicholasLiem/ModulAjar_Backend/internal/dto"
	response "github.com/NicholasLiem/ModulAjar_Backend/utils"
	"github.com/gorilla/mux"
	"net/http"
)

func (m *MicroserviceServer) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["user_id"]

	userID, err := VerifyUserId(id)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var updateUser dto.UpdateUserDTO
	err = json.NewDecoder(r.Body).Decode(&updateUser)
	if err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "Fail to update user "+err.Error())
	}

	updatedUser := datastruct.UserModel{
		UserID:   uint(userID),
		Username: updateUser.Username,
		Password: updateUser.Password,
	}

	err = m.userService.UpdateUser(updatedUser)
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, "Fail to update user "+err.Error())
		return
	}

	response.SuccessResponse(w, http.StatusOK, "Successfully updated the user", updateUser)
	return
}
