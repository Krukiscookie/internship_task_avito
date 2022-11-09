package service

import (
	"github.com/Krukiscookie/intern_task/internal/repository"
	"github.com/Krukiscookie/intern_task/pkg/models"
	"github.com/gin-gonic/gin"
)

type BalanceOperation interface {
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

type Service struct {
	BalanceOperation
	Reports
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		BalanceOperation: NewCustomerBalance(repo.CustomerBalance),
		Reports:          NewReportService(repo.Reports),
	}
}
