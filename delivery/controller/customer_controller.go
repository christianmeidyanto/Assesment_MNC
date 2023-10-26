package controller

import (
	"github.com/gin-gonic/gin"
	"mnc/merchant-bank-api/delivery/middleware"
	"mnc/merchant-bank-api/usecase"
	"net/http"
)

type CustomerController struct {
	router  *gin.Engine
	useCase usecase.CustomerUseCase
}

func (e *CustomerController) listHandler(c *gin.Context) {
	customers, err := e.useCase.FindAllCustomerList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success Get List Customer",
		"data":    customers,
	})
}
func NewCustomerController(router *gin.Engine, custUseCase usecase.CustomerUseCase) {
	controller := &CustomerController{
		router:  router,
		useCase: custUseCase,
	}
	routerGroup := controller.router.Group("/api/v1", middleware.AuthMiddleware())
	routerGroup.GET("/customer", controller.listHandler)
}
