package models

import (
	"errors"
	"net/http"
	"time"
)

type Transaction struct {
	ID              int32     `json:"transaction_id"`
	AccountID       int32     `json:"account_id"`
	OperationTypeID int32     `json:"operation_type_id"`
	Amount          float32   `json:"amount"`
	EventDate       time.Time `json:"event_date"`
}

func (t *Transaction) Bind(r *http.Request) error {
	// At this point, Decode is already done by `chi`
	if t.AccountID == 0 {
		return errors.New("invalid account_id number")
	}
	if t.OperationTypeID == 0 {
		return errors.New("invalid operationtype_id number")
	}
	if t.Amount == 0 {
		return errors.New("invalid amount number")
	}
	return nil
}
