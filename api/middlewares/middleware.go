package middlewares

import (
	"net/http"

	"caseMajoo/api/auth"
	"caseMajoo/api/models"
	"caseMajoo/api/responses"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	var response models.ResponseJson

	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			response.Success = "false"
			response.Message = "Unauthorized"
			responses.JSON(w, http.StatusUnprocessableEntity, response)
			return
		}
		next(w, r)
	}
}
