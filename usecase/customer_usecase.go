package usecase

import (
	"fmt"
	"log"
	"mnc/merchant-bank-api/model"
	"mnc/merchant-bank-api/repository"
)

type CustomerUseCase interface {
	FindAllCustomerList() ([]model.CustomerResponse, error)
	FindCustomerById(id string) (model.CustomerResponse, error)
	UpdateCustomerById(id string, amount int64) error
}

type custUseCase struct {
	repo repository.CustomerRepository
}

func (c *custUseCase) FindAllCustomerList() ([]model.CustomerResponse, error) {
	customer, err := c.repo.List()
	if err != nil {
		log.Printf("Error :%s\n", err.Error())
	}
	return customer, err
}

func (c *custUseCase) FindCustomerById(id string) (model.CustomerResponse, error) {
	return c.repo.Get(id)
}

func (c *custUseCase) UpdateCustomerById(id string, amount int64) error {
	if amount < 0 {
		return fmt.Errorf("amount not invalid")
	}
	err := c.repo.Update(id, amount)
	if err != nil {
		return fmt.Errorf("failed to update customer :%s", err.Error())
	}
	return nil
}

func NewCustomerUseCase(custRepo repository.CustomerRepository) CustomerUseCase {
	return &custUseCase{
		repo: custRepo,
	}
}
