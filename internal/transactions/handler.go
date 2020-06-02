package transactions

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/julioshinoda/api/pkg/models"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	transaction := models.Transaction{}
	if err := render.Bind(r, &transaction); err != nil {
		response, _ := json.Marshal(map[string]interface{}{"message": err.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write(response)
		return
	}
	service := NewService()
	transaction, err := service.Create(transaction)
	if err != nil {
		response, _ := json.Marshal(map[string]interface{}{"message": err.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(422)
		w.Write(response)
		return
	}
	response, _ := json.Marshal(transaction)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write(response)
}
