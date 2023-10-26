package model

import (
	"time"
)

type Transaction struct {
	Id              string    `json:"id"`
	TransactionDate time.Time `json:"transactionDate"`
	SenderId        string    `json:"SenderId" binding:"required"`
	RecieverId      string    `json:"recieverId" binding:"required"`
	Amount          int64     `json:"amount" binding:"required"`
}
