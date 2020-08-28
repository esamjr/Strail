package models

import jwt "github.com/dgrijalva/jwt-go"

//Token struct declaration
type Token struct {
	UserID   uint
	Name     string
	Username string
	*jwt.StandardClaims
}
