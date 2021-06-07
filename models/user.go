package models

import (
	"github.com/dgrijalva/jwt-go"
)

// User represents user schema
type User struct {
	Base
	Email    string `json:"email" gorm:"unique"`
	UserName string `json:"username" gorm:"unique"`
	Name     string `json:"name"`
	Number   string `json:"number" gorm:"unique"`
	Password string `json:"password"`
}

// UserErrors represent the error format for user routes
type UserErrors struct {
	Err      bool   `json:"error"`
	Email    string `json:"email"`
	Number   string `json:"number"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

// Claims represent structure of the JWT token
type Claims struct {
	jwt.StandardClaims
	ID uint `gorm:"primaryKey"`
}
