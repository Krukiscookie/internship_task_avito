package handler

import (
	"github.com/Krukiscookie/intern_task/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

// @Summary ReadFile
// @Tags reports
// @Description "Open report file"
// @Produce json
// @Param path path string true "path"
// @Success 200 {body} string response
// @Failure 500 {object} Error
// @Router /reports/{path} [get]
func (h *Handler) ReadFile(c *gin.Context) {
	filePath := c.Param("path")
	reportFile, err := os.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	c.Data(http.StatusOK, "text/csv", reportFile)
}

// @Summary TransactionInfo
// @Tags reports
// @Description "Get user transactions log"
// @Accept json
// @Produce json
// @Param input body models.GetTransactions true "JSON object with user ID, sorting method, date from and date to"
// @Success 200 {string} string response
// @Failure 500 {object} Error
// @Router /reports/transaction [post]
func (h *Handler) TransactionInfo(c *gin.Context) {
	var input models.GetTransactions

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.services.Reports.TransactionInfo(input, c)

	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"file-link": response,
	})
}

// @Summary ServiceReport
// @Tags reports
// @Description "Get monthly service report"
// @Accept json
// @Produce json
// @Param input body models.GetService true "JSON object with year and month"
// @Success 200 {string} string response
// @Failure 500 {object} Error
// @Router /reports/services-report [post]
func (h *Handler) ServiceReport(c *gin.Context) {
	var input models.GetService

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.services.Reports.ServiceReport(input, c)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"csv-file-link": response,
	})
}
