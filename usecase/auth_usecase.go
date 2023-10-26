package usecase

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"mnc/merchant-bank-api/Utils/security"
	"mnc/merchant-bank-api/model"
	"mnc/merchant-bank-api/repository"
)

type AuthUseCase interface {
	RegisterNewCustomer(payload model.Customer) (model.CustomerResponse, error)
	Login(username string, password string) (string, error)
	FindByUsernamePassword(username string, password string) (model.Auth, error)
}

type authUseCase struct {
	repoCust repository.CustomerRepository
	repo     repository.AuthRepository
}

func (a *authUseCase) RegisterNewCustomer(payload model.Customer) (model.CustomerResponse, error) {
	var customer model.CustomerResponse
	if payload.Username == "" || payload.Password == "" || payload.Name == "" || payload.Address == "" || payload.Balance == 0 {
		return customer, fmt.Errorf("username, Password, Name, Address, Balance is required")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 8)
	payload.Password = string(hashedPassword)
	customer, err = a.repoCust.Create(payload)
	if err != nil {
		return customer, fmt.Errorf("failed to create Customer: %s", err.Error())
	}
	return customer, nil
}

func (a *authUseCase) Login(username string, password string) (string, error) {
	auth, err := a.repo.GetByUsernamePassword(username, password)
	if err != nil {
		return "", fmt.Errorf("invalid username or password")
	}

	token, err := security.CreateAccessToken(auth)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %s", err.Error())
	}
	return token, nil
}

func (a *authUseCase) FindByUsernamePassword(username string, password string) (model.Auth, error) {
	return a.repo.GetByUsernamePassword(username, password)
}

func NewAuthUseCase(authRepo repository.AuthRepository, custRepo repository.CustomerRepository) AuthUseCase {
	return &authUseCase{
		repo:     authRepo,
		repoCust: custRepo,
	}
}
