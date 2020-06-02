package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julioshinoda/api/internal/accounts"
	"github.com/julioshinoda/api/internal/transactions"

	"github.com/go-chi/chi"
)

func main() {
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Println(fmt.Sprintf("Server start on PORT %s", os.Getenv("PORT")))
	http.ListenAndServe(addr, router())
}

func router() http.Handler {
	r := chi.NewRouter()

	r.Get("/accounts/{accountID}", accounts.GetByIDHandler)
	r.Post("/accounts", accounts.CreateHandler)
	r.Post("/transactions", transactions.CreateHandler)
	return r
}
