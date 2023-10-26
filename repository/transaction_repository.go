package repository

import (
	"database/sql"
	"mnc/merchant-bank-api/Utils/constant"
	"mnc/merchant-bank-api/model"
)

type TransactionRepository interface {
	Create(payload model.Transaction) (model.Transaction, error)
	Get(id string) (model.Transaction, error)
}

type transactionRepository struct {
	db *sql.DB
}

func (t *transactionRepository) Create(payload model.Transaction) (model.Transaction, error) {
	var transResponse model.Transaction
	transaction, err := t.db.Begin()
	if err != nil {
		return transResponse, err
	}

	_, err = transaction.Exec(
		constant.INSERT_TRANSACTION,
		payload.Id,
		payload.TransactionDate,
		payload.SenderId,
		payload.RecieverId,
		payload.Amount)

	if err := transaction.Commit(); err != nil {
		return transResponse, nil
	}

	return transResponse, nil
}

func (t *transactionRepository) Get(id string) (model.Transaction, error) {
	return model.Transaction{}, nil
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}
