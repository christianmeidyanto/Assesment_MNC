package model

type Auth struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
}
