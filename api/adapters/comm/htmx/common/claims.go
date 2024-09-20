package common

import (
	"github.com/golang-jwt/jwt"
)

// Claims represents the data for a user login
type Claims struct {
	UserID   string `json:"userid"`
	Roles []string `json:"roles"`
	jwt.StandardClaims
}