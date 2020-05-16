package route

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	kong "github.com/agoes/jwt-issuer/kong"
	token "github.com/agoes/jwt-issuer/token"
)

// Credentials is json request body for generating jwt
type Credentials struct {
	Username  string `json:"username"`
	Key       string `json:"key"`
	ExpiresIn int    `json:"expires_in"`
}

func createToken(w http.ResponseWriter, r *http.Request) {
	requestBody := new(Credentials)
	json.NewDecoder(r.Body).Decode(&requestBody)
	if requestBody.ExpiresIn == 0 {
		requestBody.ExpiresIn = token.DefaultExpiresIn
	}
	statusCode, _ := kong.GetConsumerJwtCredentials(requestBody.Username, requestBody.Key)
	response := make(map[string]string)
	if statusCode == 200 {
		response["token"] = token.CreateToken(requestBody.Key, requestBody.ExpiresIn)
		render.Status(r, statusCode)
	} else {
		response["error"] = "Invalid credentials"
		render.Status(r, 400)
	}
	render.JSON(w, r, response)
}

func tokenRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", createToken)
	return router
}
