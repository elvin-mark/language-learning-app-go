package handlers

import (
	"encoding/json"
	"language-learning-app/auth"
	dto "language-learning-app/dto/auth"
	models "language-learning-app/models/auth"
	"language-learning-app/services"
	"language-learning-app/utils"
	"net/http"
	"strconv"
)

type authHandlerImpl struct {
	userService services.UserService
}

// GetAuthTokenHandler godoc
//
//	@Summary		Get Auth Token
//	@Description	Given the username and password, validate it and return an authToken
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			user	body		dto.GenerateAuthTokenRequest	true	"Auth details"
//	@Success		200		{array}		models.AuthToken
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/auth/token [post]
func (h *authHandlerImpl) GetAuthTokenHandler(w http.ResponseWriter, r *http.Request) {
	var req dto.GenerateAuthTokenRequest
	var res models.AuthToken
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.WriteJSONStatus(w, map[string]string{"error": "Invalid request body"}, http.StatusBadRequest)
		return
	}
	user, err := h.userService.GetUserByUsername(req.Username)
	if err != nil || user == nil {
		utils.WriteJSONStatus(w, map[string]string{"error": "Failed authentication"}, http.StatusInternalServerError)
		return
	}
	if req.Password == user.Password {
		res.UserId = user.Id
		res.AccessToken, err = auth.GenerateToken(strconv.FormatInt(int64(res.UserId), 10), "")
		if err != nil {
			utils.WriteJSONStatus(w, map[string]string{"error": "Failed token generation"}, http.StatusInternalServerError)
			return
		}
		utils.WriteJSON(w, res)
		return
	}
	utils.WriteJSONStatus(w, map[string]string{"error": "Failed authentication"}, http.StatusInternalServerError)
}
