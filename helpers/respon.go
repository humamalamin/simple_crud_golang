package helpers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Menampilkan response dari http request
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Mengeksekusi http request
func RespondWithError(w http.ResponseWriter, code int, message string) {
	data := Response{
		Status:  int(code),
		Message: message,
		Data:    nil,
	}
	RespondWithJSON(w, code, data)
}
