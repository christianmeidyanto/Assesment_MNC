package model

import (
	"time"
)

type TransactionResponse struct {
	Id              string           `json:"id"`
	TransactionDate time.Time        `json:"transactionDate"`
	Sender          CustomerResponse `json:"SenderId" binding:"required"`
	Reciever        CustomerResponse `json:"recieverId" binding:"required"`
	Amount          int64            `json:"amount" binding:"required"`
}
