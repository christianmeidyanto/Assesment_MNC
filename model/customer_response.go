package model

type CustomerResponse struct {
	Id       string `json:"id"`
	Username string `json:"username" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Balance  int64  `json:"balance" binding:"required"`
}
