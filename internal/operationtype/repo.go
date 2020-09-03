package operationtype

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/julioshinoda/api/pkg/database"
	"github.com/julioshinoda/api/pkg/models"
)

type Repository interface {
	GetOperationTypeByID(id int32) (models.OperationType, error)
}

type OperationTypeRepo struct {
	DB *pgx.Conn
}

func NewRepository() Repository {
	conn, _ := database.GetPGConnection()
	return OperationTypeRepo{DB: conn}
}

func (repo OperationTypeRepo) GetOperationTypeByID(ID int32) (models.OperationType, error) {
	operationtype := models.OperationType{}
	query := `select ot.operationtype_id ,ot.description from operations_types ot where ot.operationtype_id  = $1`
	if err := repo.DB.QueryRow(context.Background(), query, ID).Scan(&operationtype.ID, &operationtype.Description); err != nil {
		fmt.Println("err:", err.Error())
		return operationtype, err
	}
	return operationtype, nil
}
