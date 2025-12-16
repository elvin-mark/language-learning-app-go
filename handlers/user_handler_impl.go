package handlers

import (
	"language-learning-app/services"
	"language-learning-app/utils"
	"net/http"
	"strconv"
)

type userHandlerImpl struct {
	userService services.UserService
}

// GetUserProfileHandler godoc
//
//	@Summary		Get User Profile
//	@Description	Get user profile
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		storage.User
//	@Failure		400	{object}	map[string]string
//	@Failure		500	{object}	map[string]string
//	@Router			/user/profile [get]
func (h *userHandlerImpl) GetUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("User-Id")

	if userIDStr == "" {
		utils.WriteJSONStatus(w, map[string]string{"error": "userId is required"}, http.StatusBadRequest)
		return
	}
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		utils.WriteJSONStatus(w, map[string]string{"error": "invalid userId"}, http.StatusBadRequest)
		return
	}

	user, err := h.userService.GetUserById(userID)
	if err != nil {
		utils.WriteJSONStatus(w, map[string]string{"error": "invalid userId"}, http.StatusBadRequest)
		return
	}

	utils.WriteJSON(w, user)
}
