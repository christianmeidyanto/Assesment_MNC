package model

type Customer struct {
	Id       string `json:"id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Balance  int64  `json:"balance" binding:"required"`
}
