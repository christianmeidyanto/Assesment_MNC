package repository

import (
	"database/sql"
	"fmt"
	"log"
	"mnc/merchant-bank-api/Utils/constant"
	"mnc/merchant-bank-api/model"
)

type CustomerRepository interface {
	Create(payload model.Customer) (model.CustomerResponse, error)
	List() ([]model.CustomerResponse, error)
	Get(id string) (model.CustomerResponse, error)
	Update(id string, Amount int64) error
}

type customerRepository struct {
	db *sql.DB
}

func (e *customerRepository) Create(payload model.Customer) (model.CustomerResponse, error) {
	var customer model.CustomerResponse
	_, err := e.db.Exec(constant.INSERT_CUSTOMER, payload.Id, payload.Username, payload.Password, payload.Name, payload.Address, payload.Balance)
	if err != nil {
		return customer, fmt.Errorf("failed to insert data : %s", err.Error())
	}
	return customer, nil
}

func (e *customerRepository) List() ([]model.CustomerResponse, error) {
	rows, err := e.db.Query(constant.CUSTOMER_LIST)

	if err != nil {
		return nil, err
	}

	var customers []model.CustomerResponse

	for rows.Next() {
		var customer model.CustomerResponse
		err = rows.Scan(&customer.Id, &customer.Username, &customer.Name, &customer.Address, &customer.Balance)
		if err != nil {
			log.Printf("Error scanning :%s\n", err)
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func (e *customerRepository) Get(id string) (model.CustomerResponse, error) {
	var response model.CustomerResponse
	err := e.db.QueryRow(constant.CUSTOMER_GET, id).Scan(
		&response.Id,
		&response.Username,
		&response.Name,
		&response.Address,
		&response.Balance)
	if err != nil {
		return response, fmt.Errorf("error Get Customer :%s", err.Error())
	}
	return response, nil
}

func (e *customerRepository) Update(id string, amount int64) error {
	_, err := e.db.Exec(constant.CUSTOMER_UPDATE, amount, id)
	if err != nil {
		return fmt.Errorf("error update customer :%s", err.Error())
	}
	return nil

}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRepository{
		db: db,
	}

}
