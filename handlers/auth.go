package handlers

import (
	"encoding/base64"
	"encoding/json"
	dto "language-learning-app/dto/auth"
	models "language-learning-app/models/auth"
	"language-learning-app/utils"
	"net/http"
	"os"
)

type AuthHandler interface {
	GetAuthTokenHandler(w http.ResponseWriter, r *http.Request)
}

type authHandlerImpl struct {
}

func NewAuthHandler() AuthHandler {
	return &authHandlerImpl{}
}

// ============== METHODS ==============

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
	if req.Username == os.Getenv("ADMIN_USERNAME") && req.Password == os.Getenv("ADMIN_PASSWORD") {
		res.AccessToken = base64.StdEncoding.EncodeToString([]byte(req.Username + ":" + req.Password))
		utils.WriteJSON(w, res)
		return
	}
	utils.WriteJSONStatus(w, map[string]string{"error": "Failed authentication"}, http.StatusInternalServerError)
}
