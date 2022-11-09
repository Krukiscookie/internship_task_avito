package service

import (
	"github.com/Krukiscookie/intern_task/internal/repository"
	"github.com/Krukiscookie/intern_task/pkg/models"
	"github.com/gin-gonic/gin"
)

type ReportService struct {
	repo repository.Reports
}

func NewReportService(repo repository.Reports) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) TransactionInfo(user models.GetTransactions, c *gin.Context) (string, error) {
	return s.repo.TransactionInfo(user, c)
}

func (s *ReportService) ServiceReport(user models.GetService, c *gin.Context) (string, error) {
	return s.repo.ServiceReport(user, c)
}
