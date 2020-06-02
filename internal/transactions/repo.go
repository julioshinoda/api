package transactions

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/julioshinoda/api/pkg/database"
	"github.com/julioshinoda/api/pkg/models"
)

type Repository interface {
	Create(transaction models.Transaction) (models.Transaction, error)
}

type TransactionRepo struct {
	DB *pgx.Conn
}

func NewRepository() Repository {
	conn, _ := database.GetPGConnection()
	return TransactionRepo{DB: conn}
}

func (repo TransactionRepo) Create(transaction models.Transaction) (models.Transaction, error) {
	insertSQL := `INSERT INTO transactions (account_id,operationtype_id,amount,event_date)
		VALUES ($1,$2,$3,$4) RETURNING transaction_id;`
	now := time.Now()
	eventDate := now.Format(time.RFC3339Nano)
	if err := repo.DB.QueryRow(context.Background(), insertSQL, transaction.AccountID, transaction.OperationTypeID, transaction.Amount, eventDate).Scan(&transaction.ID); err != nil {
		return transaction, err
	}
	transaction.EventDate = now
	defer repo.DB.Close(context.Background())
	return transaction, nil
}
