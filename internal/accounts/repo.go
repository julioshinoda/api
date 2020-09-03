package accounts

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/julioshinoda/api/pkg/database"
	"github.com/julioshinoda/api/pkg/models"
)

type Repository interface {
	GetByID(ID int32) (models.Account, error)
	GetByDocument(document string) (models.Account, error)
	Create(account models.Account) (models.Account, error)
}

type AccountRepo struct {
	DB *pgx.Conn
}

func NewRepository() Repository {
	conn, _ := database.GetPGConnection()
	return AccountRepo{DB: conn}
}

func (repo AccountRepo) GetByID(ID int32) (models.Account, error) {
	account := models.Account{}
	query := `select account_id,document_number from accounts where account_id = $1`
	if err := repo.DB.QueryRow(context.Background(), query, ID).Scan(&account.ID, &account.DocumentNumber); err != nil {
		fmt.Println("err:", err.Error())
		return account, err
	}
	return account, nil
}

func (repo AccountRepo) GetByDocument(document string) (models.Account, error) {
	account := models.Account{}
	query := `select account_id,document_number from accounts where document_number = $1`
	if err := repo.DB.QueryRow(context.Background(), query, document).Scan(&account.ID, &account.DocumentNumber); err != nil {
		fmt.Println("err:", err.Error())
		return account, err
	}
	return account, nil
}

func (repo AccountRepo) Create(account models.Account) (models.Account, error) {

	insertSQL := `INSERT INTO accounts (document_number)VALUES ($1) RETURNING account_id;`
	if err := repo.DB.QueryRow(context.Background(), insertSQL, account.DocumentNumber).Scan(&account.ID); err != nil {
		return account, err
	}
	return account, nil
}
