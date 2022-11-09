package handler

import (
	"errors"
	"github.com/Krukiscookie/intern_task/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary GetBalance
// @Tags user-money
// @Description "Get information about account balance"
// @Produce json
// @Param id path integer true "id"
// @Success 200 {object} models.Account
// @Failure 500 {object} Error
// @Router /user-money/{id} [get]
func (h *Handler) GetBalance(c *gin.Context) {

	idNumber, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	response, err := h.services.BalanceOperation.GetBalance(models.Account{Id: int(idNumber)})

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"balance": response.Balance,
		"reserve": response.Reserve,
	})
}

// @Summary TransferMoney
// @Tags user-money
// @Description "Transferring money to another user"
// @Accept json
// @Produce json
// @Param input  body models.SwagTransaction true "JSON object with ID_from, ID_to, money amount and status to transfer"
// @Success 200 {string} string "successfully transfer money"
// @Failure 500 {object} Error
// @Router /user-money/transfer [post]
func (h *Handler) TransferMoney(c *gin.Context) {
	var input models.Transaction

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response, _ := h.services.BalanceOperation.GetBalance(models.Account{Id: input.IdFrom})

	if input.Amount > response.Balance {
		err := errors.New("not enough money")
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	} else if input.Amount <= 0 {
		err := errors.New("sum less or equal 0")
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.TransferMoney(input)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully transfer money",
	})
}

// @Summary AddMoney
// @Tags user-money
// @Description "Add money for a given account"
// @Accept json
// @Produce json
// @Param input body models.SwagAccount true "JSON with user ID and money amount"
// @Success 200 {string} string "successfully add money"
// @Failure 500 {object} Error
// @Router /user-money/addmoney [post]
func (h *Handler) AddMoney(c *gin.Context) {
	var input models.Account

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.Id <= 0 || input.Balance < 0 {
		err := errors.New("id, sum less or equal 0")
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.AddMoney(input)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully add money",
	})
}

// @Summary ServiceReserve
// @Tags user-money
// @Description "Сreates an order and reserves money on the account"
// @Accept json
// @Produce json
// @Param input body models.SwagServices true "JSON object with user ID, service ID, order ID and amount"
// @Success 200 {string} string "successfully reserve money for payment"
// @Failure 500 {object} Error
// @Router /user-money/services/reserve [post]
func (h *Handler) ServiceReserve(c *gin.Context) {
	var input models.Services

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.AccountId <= 0 || input.Amount <= 0 {
		err := errors.New("id, sum less or equal 0")
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response, _ := h.services.BalanceOperation.GetBalance(models.Account{Id: input.AccountId})

	if input.Amount > response.Balance {
		err := errors.New("not enough money")
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.ServiceReserve(input)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully reserve money for payment",
	})
}

// @Summary ServiceApprove
// @Tags user-money
// @Description "Confirms the payment and writes off the money from the reserve"
// @Accept json
// @Produce json
// @Param input body models.SwagServices true "JSON object with user ID, service ID, order ID and amount"
// @Success 200 {string} string "successful payment"
// @Failure 500 {object} Error
// @Router /user-money/services/approve [post]
func (h *Handler) ServiceApprove(c *gin.Context) {
	var input models.Services

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.AccountId <= 0 || input.Amount <= 0 {
		err := errors.New("id, sum less or equal 0")
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.ServiceApprove(input)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successful payment",
	})
}

// @Summary ServiceRefusal
// @Tags user-money
// @Description "Cancels the payment and returns the money to the balance"
// @Accept json
// @Produce json
// @Param input body models.SwagServices true "JSON object with user ID, service ID, order ID and amount"
// @Success 200 {string} string "payment cancellation"
// @Failure 500 {object} Error
// @Router /user-money/services/refusal [post]
func (h *Handler) ServiceRefusal(c *gin.Context) {
	var input models.Services

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.AccountId <= 0 || input.Amount <= 0 {
		err := errors.New("id, sum less or equal 0")
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.ServiceRefusal(input)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "payment cancellation",
	})
}
