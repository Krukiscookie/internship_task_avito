package repository

import (
	"errors"
	"fmt"
	"github.com/Krukiscookie/intern_task/pkg/models"
	"github.com/jmoiron/sqlx"
	"strconv"
	"time"
)

type BalancePostgres struct {
	db *sqlx.DB
}

func NewBalancePostgres(db *sqlx.DB) *BalancePostgres {
	return &BalancePostgres{db: db}
}

func (r *BalancePostgres) NewAccount(account models.Account) error {
	query := fmt.Sprintf("INSERT INTO %s (id, balance) VALUES ($1, 0)", accountsTable)

	if _, err := r.db.Query(query, account.Id); err != nil {
		return err
	}

	return nil
}

func (r *BalancePostgres) GetBalance(account models.Account) (models.Account, error) {
	var balance, reserve float64

	query := fmt.Sprintf("SELECT balance, reserve FROM %s WHERE id = $1", accountsTable)
	row := r.db.QueryRow(query, account.Id)

	if err := row.Scan(&balance, &reserve); err != nil {
		return models.Account{}, err
	}

	return models.Account{Balance: balance, Reserve: reserve}, nil
}

func (r *BalancePostgres) AddMoney(account models.Account) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	updateQuery := fmt.Sprintf("UPDATE %s SET balance = balance + $2 WHERE id = $1", accountsTable)

	if _, err := tx.Exec(updateQuery, account.Id, account.Balance); err != nil {
		tx.Rollback()
		return err
	}

	insertQuery := fmt.Sprintf("INSERT INTO %s (id_from, amount, status, create_time) VALUES ($1, $2, $3, $4)", transactionTable)

	_, err = tx.Exec(insertQuery, account.Id, account.Balance, "refill", time.Now())
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *BalancePostgres) TransferMoney(transact models.Transaction) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	updateQueryFrom := fmt.Sprintf("UPDATE %s SET balance = balance - $2 WHERE id = $1", accountsTable)

	if _, err := tx.Exec(updateQueryFrom, transact.IdFrom, transact.Amount); err != nil {
		tx.Rollback()
		return err
	}

	updateQueryTo := fmt.Sprintf("UPDATE %s SET balance = balance + $2 WHERE id = $1", accountsTable)

	if _, err := tx.Exec(updateQueryTo, transact.IdTo, transact.Amount); err != nil {
		tx.Rollback()
		return err
	}

	insertQuery := fmt.Sprintf("INSERT INTO %s (id_from, id_to, amount, status, create_time) VALUES ($1, $2, $3, $4, $5)", transactionTable)

	_, err = tx.Exec(insertQuery, transact.IdFrom, transact.IdTo, transact.Amount, transact.Status, time.Now())
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *BalancePostgres) ServiceReserve(service models.Services) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	userQuery := fmt.Sprintf("UPDATE %s SET reserve = reserve + $2, balance = balance - $2 WHERE id = $1", accountsTable)

	if _, err := tx.Exec(userQuery, service.AccountId, service.Amount); err != nil {
		tx.Rollback()
		return err
	}

	serviceQuery := fmt.Sprintf("INSERT INTO %s (account_id, amount, service_id, order_id, status, created_at) VALUES ($1, $2, $3, $4, $5, $6)", servicesTable)

	_, err = tx.Exec(serviceQuery, service.AccountId, service.Amount, service.IdService, service.IdOrder, "Pending", time.Now())
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *BalancePostgres) ServiceApprove(service models.Services) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	serviceQuery := fmt.Sprintf("UPDATE %s SET status=$6, updated_at=$7 WHERE account_id=$1 AND service_id=$2 AND order_id=$3 AND amount=$4 AND status=$5", servicesTable)

	res, err := tx.Exec(serviceQuery, service.AccountId, service.IdService, service.IdOrder, service.Amount, "Pending", "Success", time.Now())
	rows, _ := res.RowsAffected()

	if rows == 0 {
		tx.Rollback()
		return errors.New("order doesn't exist or already approved")
	} else if err != nil {
		tx.Rollback()
		return err
	}

	userQuery := fmt.Sprintf("UPDATE %s SET reserve = reserve - $2 WHERE id = $1", accountsTable)

	if _, err := tx.Exec(userQuery, service.AccountId, service.Amount); err != nil {
		tx.Rollback()
		return err
	}

	transactQuery := fmt.Sprintf("INSERT INTO %s (id_from, amount, status, create_time) VALUES ($1, $2, $3, $4)", transactionTable)

	status := "Payment for service:" + strconv.Itoa(service.IdService) + ", order:" + strconv.Itoa(service.IdOrder)
	_, err = tx.Exec(transactQuery, service.AccountId, service.Amount, status, time.Now())
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *BalancePostgres) ServiceRefusal(service models.Services) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	serviceQuery := fmt.Sprintf("UPDATE %s SET status = $6, updated_at = $7 WHERE account_id=$1 AND service_id=$2 AND order_id=$3 AND amount=$4 AND status=$5", servicesTable)

	res, err := tx.Exec(serviceQuery, service.AccountId, service.IdService, service.IdOrder, service.Amount, "Pending", "Cancel", time.Now())
	rows, _ := res.RowsAffected()

	if rows == 0 {
		tx.Rollback()
		return errors.New("order doesn't exist or already canceled")
	} else if err != nil {
		tx.Rollback()
		return err
	}

	userQuery := fmt.Sprintf("UPDATE %s SET balance = balance + $2, reserve = reserve - $2 WHERE id = $1", accountsTable)

	if _, err := tx.Exec(userQuery, service.AccountId, service.Amount); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
