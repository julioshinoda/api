package transactions

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/julioshinoda/api/internal/accounts"
	"github.com/julioshinoda/api/internal/operationtype"
	accountMock "github.com/julioshinoda/api/mocks/accounts"
	operationMock "github.com/julioshinoda/api/mocks/operationtype"
	transactionMock "github.com/julioshinoda/api/mocks/transactions"
	"github.com/julioshinoda/api/pkg/models"
)

func TestIsNegativeValue(t *testing.T) {
	type args struct {
		description string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Is negative value",
			args: args{
				description: "Compra",
			},
			want: true,
		},
		{
			name: "Is not negative value",
			args: args{
				description: "pagamento",
			},
			want: false,
		},
	}
	os.Setenv("NEGATIVE_TYPES", "compra,saque")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNegativeValue(tt.args.description); got != tt.want {
				t.Errorf("IsNegativeValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransactionService_Create(t *testing.T) {
	type fields struct {
		Repo              Repository
		OperationTypeRepo operationtype.Repository
		AccountRepo       accounts.Repository
	}
	type args struct {
		transaction models.Transaction
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.Transaction
		wantErr error
	}{
		{
			name: "success to create transaction with type pagamento",
			fields: fields{
				Repo:              &transactionMock.Repository{},
				AccountRepo:       &accountMock.Repository{},
				OperationTypeRepo: &operationMock.Repository{},
			},
			args: args{
				transaction: models.Transaction{
					AccountID:       int32(1),
					OperationTypeID: int32(1),
					Amount:          12.50,
				},
			},
			want: models.Transaction{
				ID:              int32(1),
				AccountID:       int32(1),
				OperationTypeID: int32(1),
				Amount:          12.50,
			},
		},
		{
			name: "success to create transaction with type saque",
			fields: fields{
				Repo:              &transactionMock.Repository{},
				AccountRepo:       &accountMock.Repository{},
				OperationTypeRepo: &operationMock.Repository{},
			},
			args: args{
				transaction: models.Transaction{
					AccountID:       int32(1),
					OperationTypeID: int32(1),
					Amount:          12.50,
				},
			},
			want: models.Transaction{
				ID:              int32(1),
				AccountID:       int32(1),
				OperationTypeID: int32(1),
				Amount:          -12.50,
			},
		},
		{
			name: "operation type not found",
			fields: fields{
				Repo:              &transactionMock.Repository{},
				AccountRepo:       &accountMock.Repository{},
				OperationTypeRepo: &operationMock.Repository{},
			},
			args: args{
				transaction: models.Transaction{
					AccountID:       int32(1),
					OperationTypeID: int32(1),
					Amount:          12.50,
				},
			},
			want:    models.Transaction{},
			wantErr: fmt.Errorf("not found operationtype_id %v", int32(1)),
		},
		{
			name: "account not found",
			fields: fields{
				Repo:              &transactionMock.Repository{},
				AccountRepo:       &accountMock.Repository{},
				OperationTypeRepo: &operationMock.Repository{},
			},
			args: args{
				transaction: models.Transaction{
					AccountID:       int32(1),
					OperationTypeID: int32(1),
					Amount:          12.50,
				},
			},
			want:    models.Transaction{},
			wantErr: fmt.Errorf("not found account_id %v", int32(1)),
		},
	}
	os.Setenv("NEGATIVE_TYPES", "compra,saQue")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := TransactionService{
				Repo:              tt.fields.Repo,
				OperationTypeRepo: tt.fields.OperationTypeRepo,
				AccountRepo:       tt.fields.AccountRepo,
			}
			switch tt.name {
			case "success to create transaction with type pagamento":
				service.OperationTypeRepo.(*operationMock.Repository).On("GetOperationTypeByID", tt.args.transaction.OperationTypeID).Return(models.OperationType{ID: 1, Description: "PAGAMENTO"}, nil).Once()
				service.AccountRepo.(*accountMock.Repository).On("GetByID", tt.args.transaction.AccountID).Return(models.Account{}, nil).Once()
				response := models.Transaction{
					ID:              1,
					AccountID:       1,
					OperationTypeID: 1,
					Amount:          12.50,
				}
				service.Repo.(*transactionMock.Repository).On("Create", tt.args.transaction).Return(response, nil).Once()
			case "success to create transaction with type saque":
				service.OperationTypeRepo.(*operationMock.Repository).On("GetOperationTypeByID", tt.args.transaction.OperationTypeID).Return(models.OperationType{ID: 1, Description: "saque"}, nil).Once()
				service.AccountRepo.(*accountMock.Repository).On("GetByID", tt.args.transaction.AccountID).Return(models.Account{}, nil).Once()
				param := models.Transaction{
					AccountID:       1,
					OperationTypeID: 1,
					Amount:          -12.50,
				}

				service.Repo.(*transactionMock.Repository).On("Create", param).Return(tt.want, nil).Once()
			case "operation type not found":
				service.OperationTypeRepo.(*operationMock.Repository).On("GetOperationTypeByID", tt.args.transaction.OperationTypeID).Return(models.OperationType{}, errors.New("no rows in result set")).Once()

			case "account not found":
				service.OperationTypeRepo.(*operationMock.Repository).On("GetOperationTypeByID", tt.args.transaction.OperationTypeID).Return(models.OperationType{ID: 1, Description: "saque"}, nil).Once()
				service.AccountRepo.(*accountMock.Repository).On("GetByID", tt.args.transaction.AccountID).Return(models.Account{}, errors.New("no rows in result set")).Once()
			}
			got, err := service.Create(tt.args.transaction)
			if (err != nil) && tt.wantErr.Error() != err.Error() {
				t.Errorf("TransactionService.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransactionService.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
