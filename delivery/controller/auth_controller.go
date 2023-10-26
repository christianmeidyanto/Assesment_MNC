package controller

import (
	"github.com/gin-gonic/gin"
	"mnc/merchant-bank-api/Utils/common"
	"mnc/merchant-bank-api/model"
	"mnc/merchant-bank-api/usecase"
	"net/http"
)

type AuthController struct {
	router   *gin.Engine
	authCase usecase.AuthUseCase
}

func (a *AuthController) createHandler(c *gin.Context) {
	var customer model.Customer
	var response model.CustomerResponse
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	customer.Id = common.GenerateUUID()
	_, err := a.authCase.RegisterNewCustomer(customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response.Id = customer.Id
	response.Username = customer.Username
	response.Name = customer.Name
	response.Address = customer.Address
	response.Balance = customer.Balance

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Success Create New Customer",
		"data":    response,
	})
}
func (a *AuthController) loginHandler(c *gin.Context) {
	var payload model.Auth
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := a.authCase.Login(payload.Username, payload.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sign In Success",
		"token":   token,
	})
	return
}

func NewAuthController(route *gin.Engine, authUseCase usecase.AuthUseCase) {
	controller := AuthController{
		router:   route,
		authCase: authUseCase,
	}

	routerGroup := controller.router.Group("/api/v1")
	routerGroup.POST("/register", controller.createHandler)
	routerGroup.POST("/login", controller.loginHandler)
}
