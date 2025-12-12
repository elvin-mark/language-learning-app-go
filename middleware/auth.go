package middleware

import (
	"language-learning-app/services"
	"net/http"
	"strconv"
)

func NewBasicAuth(userService services.UserService) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			username, password, ok := r.BasicAuth()
			user, err := userService.GetUserByUsername(username)
			if !ok || err != nil || user == nil || user.Password != password {
				w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			r.Header.Set("User-Id", strconv.FormatInt(int64(user.UserID), 10))
			next.ServeHTTP(w, r)
		})
	}
}
