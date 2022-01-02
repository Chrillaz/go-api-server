package handlers

import (
	"encoding/json"
	"net/http"

	"go-api-server/api/mocks"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Todos)
}
