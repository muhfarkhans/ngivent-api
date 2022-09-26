package handler

import (
	"errors"
	"fmt"
	"net/http"
	"ngevent-api/helper"
	"ngevent-api/transaction"
	"ngevent-api/user"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) CreateNewTransaction(c *gin.Context) {
	var input transaction.CreateTransactionInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errorValidation := helper.FormatValidationError(err)
			errors := gin.H{"errors": errorValidation}

			response := helper.APIResponse("error validation", http.StatusUnprocessableEntity, "error", errors)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}

		errors := gin.H{"errors": err.Error()}
		response := helper.APIResponse("error create transaction", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	OrderId := fmt.Sprintf("%d%d", currentUser.Id, time.Now().UnixNano())
	input.OrderId = OrderId

	newTransaction, err := h.service.CreateTransaction(input)
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := helper.APIResponse("error create transaction", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := transaction.FormatTransaction(newTransaction)
	response := helper.APIResponse("success create transaction", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetAllTransaction(c *gin.Context) {
	param := helper.Pagination{}
	c.BindQuery(&param)

	currentUser := c.MustGet("currentUser").(user.User)
	allTransaction, paginate, err := h.service.GetTransactions(param, currentUser)
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := helper.APIResponse("error get transaction", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := transaction.FormatTransactions(allTransaction, paginate)

	response := helper.APIResponse("success get transactions", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetTransaction(c *gin.Context) {
	var input transaction.GetTransactionDetailInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errorValidation := helper.FormatValidationError(err)
			errors := gin.H{"errors": errorValidation}

			response := helper.APIResponse("error validation", http.StatusUnprocessableEntity, "error", errors)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}

		errors := gin.H{"errors": err.Error()}
		response := helper.APIResponse("error create transaction", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	detailTransaction, err := h.service.GetTransaction(input, currentUser)
	if err != nil {
		errors := gin.H{"errors": err.Error()}

		response := helper.APIResponse("error get transaction", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := transaction.FormatDetailTransaction(detailTransaction)

	response := helper.APIResponse("success get transactions", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetNotification(c *gin.Context) {
	var input transaction.TransactionNotificationInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create transaction", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = h.service.ProcessPayment(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Failed to create transaction", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, input)
}
