package accounts

import (
	"fmt"
	"strconv"

	"github.com/julioshinoda/api/pkg/models"
)

type Service interface {
	GetByID(ID string) (models.Account, error)
	Create(account models.Account) (models.Account, error)
}

type AccountService struct {
	Repo Repository
}

func NewService() Service {
	return AccountService{Repo: NewRepository()}
}

func (service AccountService) GetByID(ID string) (models.Account, error) {
	accountID, err := strconv.Atoi(ID)
	if err != nil {
		return models.Account{}, err
	}
	return service.Repo.GetByID(int32(accountID))
}

func (service AccountService) Create(account models.Account) (models.Account, error) {
	if _, err := service.Repo.GetByDocument(account.DocumentNumber); err == nil {
		return models.Account{}, fmt.Errorf("Document %s already created", account.DocumentNumber)
	}
	return service.Repo.Create(account)
}
