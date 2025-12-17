package handlers

import (
	"encoding/json"
	dto "language-learning-app/dto/user"
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
//	@Success		200	{object}	storage.User
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

// UpdateUserSettingsHandler godoc
//
//	@Summary		Update user profile settings
//	@Description	Update user profile settings
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user	body		dto.UpdateUserSettings	true	"New User settings"
//	@Success		200		{object}	storage.User
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/user/profile [patch]
func (h *userHandlerImpl) UpdateUserSettingsHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateUserSettings
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteJSONStatus(w, map[string]string{"error": "Invalid request body"}, http.StatusBadRequest)
		return
	}

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

	user, err := h.userService.UpdateUserSettings(userID, req.PreferredLanguage, req.TargetLanguage)
	if err != nil {
		utils.WriteJSONStatus(w, map[string]string{"error": "could not update user settings"}, http.StatusBadRequest)
		return
	}
	utils.WriteJSON(w, user)
}
