package models

import (
	"errors"
	"net/http"
)

type Account struct {
	ID             int32  `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

func (a *Account) Bind(r *http.Request) error {
	// At this point, Decode is already done by `chi`
	if len(a.DocumentNumber) != 11 {
		return errors.New("invalid document number")
	}

	return nil
}
