package response

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

type BaseReponse struct {
	Data    *interface{} `json:"data"json:"data,omitempty"`
	Error   *string      `json:"error,omitempty"`
	Message *string      `json:"message,omitempty"`
}

func ResponseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	respond(w, code, BaseReponse{Data: &payload})
}

func ResponseWithError(w http.ResponseWriter, code int, errorMessage string) {
	respond(w, code, BaseReponse{Error: &errorMessage})
}

func respond(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		log.Error().Err(err).Msg("Error on Return Response")
	}

}
