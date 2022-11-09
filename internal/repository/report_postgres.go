package repository

import (
	"fmt"
	"github.com/Krukiscookie/intern_task/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joho/sqltocsv"
	"os"
	"strconv"
	"strings"
	"time"
)

type ReportsPostgres struct {
	db *sqlx.DB
}

func NewReportsPostgres(db *sqlx.DB) *ReportsPostgres {
	return &ReportsPostgres{db: db}
}

func (r *ReportsPostgres) TransactionInfo(user models.GetTransactions, c *gin.Context) (string, error) {
	var param, order string

	switch user.SortBy {
	case "date":
		param = "create_time"
	case "amount":
		order = "amount"
	}
	switch user.SortOrder {
	case "descending":
		order = "DESC"
	case "ascending":
		order = "ASC"
	}

	getTransactionByUserFromViewQuery := fmt.Sprintf("SELECT * FROM %s WHERE id_from = $1 AND DATE (create_time) BETWEEN $2 AND $3 ORDER BY $4 LIMIT 20 OFFSET $5", transactionTable)

	rows, err := r.db.QueryContext(
		c,
		getTransactionByUserFromViewQuery,
		user.User,
		user.DateFrom,
		user.DateTo,
		param+" "+order,
		(user.NumberOfPage-1)*20,
	)

	if err != nil {
		return "", err
	}

	fileName, err := CreateFile("user " + strconv.Itoa(int(user.User)) + " report ")
	if err != nil {
		return "", err
	}
	err = sqltocsv.WriteFile(fileName, rows)
	if err != nil {
		panic(err)
	}

	return fileName, nil
}

func (r *ReportsPostgres) ServiceReport(user models.GetService, c *gin.Context) (string, error) {

	date := user.Year + "-" + user.Month + "-" + "01"
	query := fmt.Sprintf("SELECT service_id, SUM(amount) FROM %s WHERE status = 'Success' AND date_trunc('month', updated_at)::date = '%s'::date GROUP BY service_id", servicesTable, date)

	rows, err := r.db.QueryContext(c, query)

	fileName, newErr := CreateFile("services_")
	if newErr != nil {
		return "", err
	}
	err = sqltocsv.WriteFile(fileName, rows)
	if err != nil {
		panic(err)
	}

	return fileName, err
}

func CreateFile(filePrefix string) (string, error) {

	fileName := strings.ReplaceAll(filePrefix+time.Now().Format(time.RFC822)+".csv", " ", "_")

	file, err := os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return "", err
	}
	err = file.Close()
	if err != nil {
		return "", err
	}

	return fileName, nil
}
