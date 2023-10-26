package delivery

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"mnc/merchant-bank-api/config"
	"mnc/merchant-bank-api/delivery/controller"
	"mnc/merchant-bank-api/repository"
	"mnc/merchant-bank-api/usecase"
)

type appServer struct {
	custUseCase  usecase.CustomerUseCase
	transUseCase usecase.TransactionUseCase
	authUseCase  usecase.AuthUseCase
	engine       *gin.Engine
	host         string
}

func (a *appServer) initController() {
	controller.NewCustomerController(a.engine, a.custUseCase)
	controller.NewTransactionController(a.engine, a.transUseCase)
	controller.NewAuthController(a.engine, a.authUseCase)
}

func (a *appServer) Run() {
	a.initController()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err.Error())
	}
}

func Server() *appServer {
	engine := gin.Default()
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	dbConn, _ := config.NewDbConnection(cfg)
	custRepo := repository.NewCustomerRepository(dbConn.Conn())
	custUseCase := usecase.NewCustomerUseCase(custRepo)
	authRepo := repository.NewAuthRepository(dbConn.Conn())
	authUseCase := usecase.NewAuthUseCase(authRepo, custRepo)
	transRepo := repository.NewTransactionRepository(dbConn.Conn())
	transUseCase := usecase.NewTransactionUseCase(transRepo, custUseCase)
	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	return &appServer{
		engine:       engine,
		custUseCase:  custUseCase,
		authUseCase:  authUseCase,
		transUseCase: transUseCase,
		host:         host,
	}
}
