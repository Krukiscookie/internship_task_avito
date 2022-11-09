package service

import (
	"database/sql"
	"github.com/Krukiscookie/intern_task/internal/repository"
	"github.com/Krukiscookie/intern_task/pkg/models"
)

type CustomerBalance struct {
	repo repository.CustomerBalance
}

func NewCustomerBalance(repo repository.CustomerBalance) *CustomerBalance {
	return &CustomerBalance{repo: repo}
}

func (s *CustomerBalance) GetBalance(account models.Account) (models.Account, error) {
	return s.repo.GetBalance(account)
}

func (s *CustomerBalance) AddMoney(account models.Account) error {
	_, err := s.repo.GetBalance(account)
	if err != nil && err == sql.ErrNoRows {
		if err := s.repo.NewAccount(account); err != nil {
			return err
		}
	}

	return s.repo.AddMoney(account)
}

func (s *CustomerBalance) TransferMoney(transact models.Transaction) error {
	return s.repo.TransferMoney(transact)
}

func (s *CustomerBalance) ServiceReserve(service models.Services) error {
	return s.repo.ServiceReserve(service)
}

func (s *CustomerBalance) ServiceApprove(service models.Services) error {
	return s.repo.ServiceApprove(service)
}

func (s *CustomerBalance) ServiceRefusal(service models.Services) error {
	return s.repo.ServiceRefusal(service)
}
