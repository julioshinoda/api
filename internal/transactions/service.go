package transactions

import (
	"fmt"
	"os"
	"strings"

	"github.com/julioshinoda/api/internal/accounts"
	"github.com/julioshinoda/api/internal/operationtype"
	"github.com/julioshinoda/api/pkg/models"
)

type Service interface {
	Create(transaction models.Transaction) (models.Transaction, error)
}

type TransactionService struct {
	Repo              Repository
	OperationTypeRepo operationtype.Repository
	AccountRepo       accounts.Repository
}

func NewService() Service {
	return TransactionService{
		Repo:              NewRepository(),
		OperationTypeRepo: operationtype.NewRepository(),
		AccountRepo:       accounts.NewRepository(),
	}
}

func (service TransactionService) Create(transaction models.Transaction) (models.Transaction, error) {
	operationType, err := service.OperationTypeRepo.GetOperationTypeByID(transaction.OperationTypeID)
	if err != nil {
		return models.Transaction{}, fmt.Errorf("not found operationtype_id %v", transaction.OperationTypeID)
	}
	if _, err := service.AccountRepo.GetByID(transaction.AccountID); err != nil {
		return models.Transaction{}, fmt.Errorf("not found account_id %v", transaction.AccountID)
	}
	if IsNegativeValue(operationType.Description) {
		transaction.Amount *= -1
	}
	return service.Repo.Create(transaction)
}

func IsNegativeValue(description string) bool {
	negativeType := os.Getenv("NEGATIVE_TYPES")
	types := strings.Split(negativeType, ",")
	for _, t := range types {
		if strings.Contains(strings.ToLower(description), strings.ToLower(t)) {
			return true
		}
	}
	return false
}
