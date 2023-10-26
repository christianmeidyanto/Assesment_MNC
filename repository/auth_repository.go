package repository

import (
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"mnc/merchant-bank-api/Utils/constant"
	"mnc/merchant-bank-api/model"
)

type AuthRepository interface {
	Create(payload model.Customer) (model.CustomerResponse, error)
	GetByUsername(username string) (model.Auth, error)
	GetByUsernamePassword(username string, password string) (model.Auth, error)
}

type authRepository struct {
	db *sql.DB
}

func (e *authRepository) Create(payload model.Customer) (model.CustomerResponse, error) {
	var customer model.CustomerResponse
	_, err := e.db.Exec(constant.INSERT_CUSTOMER, payload.Id, payload.Username, payload.Password, payload.Name, payload.Address, payload.Balance)
	if err != nil {
		return customer, fmt.Errorf("failed to insert data : %s", err.Error())
	}
	return customer, nil
}

func (e *authRepository) GetByUsername(username string) (model.Auth, error) {
	var auth model.Auth
	err := e.db.QueryRow(constant.CUSTOMER_GET_USER, username).Scan(
		&auth.Username,
		&auth.Password,
	)
	if err != nil {
		return model.Auth{}, fmt.Errorf("error Get Username Customer :%s", err.Error())
	}
	return auth, nil
}

func (e *authRepository) GetByUsernamePassword(username string, password string) (model.Auth, error) {
	auth, err := e.GetByUsername(username)
	if err != nil {
		return model.Auth{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(auth.Password), []byte(password))
	if err != nil {
		return model.Auth{}, err
	}
	return auth, nil
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{
		db: db,
	}

}
