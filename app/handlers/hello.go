package handlers

import (
	"encoding/json"
	"net/http"
)

type HelloResponse struct {
	Message string `json:"message"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	response := HelloResponse{
		Message: "GoodBYE!!!!",
	}
	json.NewEncoder(w).Encode(response)
	return
}
