package controller

import (
	"github.com/gin-gonic/gin"
	"mnc/merchant-bank-api/Utils/common"
	"mnc/merchant-bank-api/delivery/middleware"
	"mnc/merchant-bank-api/model"
	"mnc/merchant-bank-api/usecase"
	"net/http"
)

type TransactionController struct {
	router             *gin.Engine
	transactionUseCase usecase.TransactionUseCase
}

func (t *TransactionController) createHandler(c *gin.Context) {
	var transaction model.Transaction
	err := c.ShouldBindJSON(&transaction)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	transaction.Id = common.GenerateUUID()
	transactionResponse, err := t.transactionUseCase.RegisterNewTransaction(transaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Success Create New Transaction",
		"data":    transactionResponse,
	})
}

func NewTransactionController(t *gin.Engine, transUseCase usecase.TransactionUseCase) {
	controller := TransactionController{
		router:             t,
		transactionUseCase: transUseCase,
	}
	routerGroup := controller.router.Group("/api/v1", middleware.AuthMiddleware())
	routerGroup.POST("/transaction", controller.createHandler)
}
