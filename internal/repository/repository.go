package repository

import (
	"github.com/Krukiscookie/intern_task/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type CustomerBalance interface {
	NewAccount(account models.Account) error
	GetBalance(account models.Account) (models.Account, error)
	AddMoney(account models.Account) error
	TransferMoney(transact models.Transaction) error
	ServiceReserve(service models.Services) error
	ServiceApprove(service models.Services) error
	ServiceRefusal(service models.Services) error
}

type Reports interface {
	TransactionInfo(user models.GetTransactions, c *gin.Context) (string, error)
	ServiceReport(user models.GetService, c *gin.Context) (string, error)
}

type Repository struct {
	CustomerBalance
	Reports
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		CustomerBalance: NewBalancePostgres(db),
		Reports:         NewReportsPostgres(db),
	}
}
