package common

import (
	"github.com/golang-jwt/jwt"
)

// Claims represents the data for a user login
type Claims struct {
	UserID   string `json:"userid"`
	Roles []string `json:"roles"`
	UserName string `json:"username"`
	jwt.StandardClaims
}

// Link represents a link in the navigation bar
type Link struct {
	URL string `json:"url"`
	Text string `json:"text"`
}

