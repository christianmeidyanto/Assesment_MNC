package usecase

import (
	"fmt"
	"log"
	"mnc/merchant-bank-api/model"
	"mnc/merchant-bank-api/repository"
	"time"
)

type TransactionUseCase interface {
	RegisterNewTransaction(payload model.Transaction) (model.TransactionResponse, error)
}

type transactionUseCase struct {
	repo        repository.TransactionRepository
	custUseCase CustomerUseCase
}

func (t *transactionUseCase) RegisterNewTransaction(payload model.Transaction) (model.TransactionResponse, error) {
	var transactionResponse model.TransactionResponse
	var transaction model.Transaction
	// get customer sender
	sender, err := t.custUseCase.FindCustomerById(payload.SenderId)

	if err != nil {
		return transactionResponse, fmt.Errorf("sender not found")
	}

	// get customer reciever
	reciever, err := t.custUseCase.FindCustomerById(payload.RecieverId)
	log.Println(reciever)
	if err != nil {
		return transactionResponse, fmt.Errorf("reciever not found")
	}

	// check balance sender
	if payload.Amount > sender.Balance {
		return transactionResponse, fmt.Errorf("your balance is low")
	}

	now := time.Now()
	format := now.Format("02-01-2006 15:04:05")
	parse, err := time.Parse("02-01-2006 15:04:05", format)
	payload.TransactionDate = parse
	transaction, err = t.repo.Create(payload)
	if err != nil {
		return transactionResponse, fmt.Errorf("transaction Failed :%s", err.Error())
	}
	err = t.custUseCase.UpdateCustomerById(reciever.Id, reciever.Balance+payload.Amount)
	if err != nil {
		return transactionResponse, fmt.Errorf("transaction Failed :%s", err.Error())
	}
	err = t.custUseCase.UpdateCustomerById(sender.Id, sender.Balance-payload.Amount)
	if err != nil {
		return transactionResponse, fmt.Errorf("transaction Failed :%s", err.Error())
	}
	sender, err = t.custUseCase.FindCustomerById(payload.SenderId)
	reciever, err = t.custUseCase.FindCustomerById(payload.RecieverId)

	transactionResponse.Id = payload.Id
	transactionResponse.TransactionDate = transaction.TransactionDate
	transactionResponse.Sender = sender
	transactionResponse.Reciever = reciever
	transactionResponse.Amount = payload.Amount
	return transactionResponse, nil
}

func NewTransactionUseCase(repository repository.TransactionRepository, customerUseCase CustomerUseCase) TransactionUseCase {
	return &transactionUseCase{
		repo:        repository,
		custUseCase: customerUseCase,
	}
}
