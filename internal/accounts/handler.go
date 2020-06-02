package accounts

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/julioshinoda/api/pkg/models"
)

func GetByIDHandler(w http.ResponseWriter, r *http.Request) {
	accountIDParam := chi.URLParam(r, "accountID")
	service := NewService()
	account, err := service.GetByID(accountIDParam)
	if err != nil {
		response, _ := json.Marshal(map[string]interface{}{"message": err.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(422)
		w.Write(response)
		return
	}
	response, _ := json.Marshal(account)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(response)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	account := models.Account{}
	if err := render.Bind(r, &account); err != nil {
		response, _ := json.Marshal(map[string]interface{}{"message": err.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write(response)
		return
	}
	service := NewService()
	account, err := service.Create(account)
	if err != nil {
		response, _ := json.Marshal(map[string]interface{}{"message": err.Error()})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(422)
		w.Write(response)
		return
	}
	response, _ := json.Marshal(account)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write(response)
}
