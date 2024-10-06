package middleware

import (
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/serdarkalayci/membership/api/adapters/comm/htmx/common"
)
const secretKey = "the_most_secure_secret"
const cookieName = "membershiptoken"

// Authenticate is a middleware that checks if the request is authenticated
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := CheckAuthentication(c)
		if err != nil {
			log.Err(err).Msg("Authentication failed")
			c.Header("HX-Redirect", "/")
			return
		}
		c.Set("UserID", claims.UserID)
		c.Set("Roles", claims.Roles)
		c.Set("UserName", claims.UserName)
		c.Next()
	}
}

// CheckAuthentication is a function that checks if the request is authenticated
func CheckAuthentication(c *gin.Context) (*common.Claims, error) {
	tknStr, err := c.Cookie(cookieName)
	if err != nil {
		if err == http.ErrNoCookie {
			return nil, common.ErrCookieNotFound{}
		}
		return nil, err
	}

	// Get the JWT string from the cookie
	var jwtKey = []byte(secretKey)
	// Initialize a new instance of `Claims`
	claims := &common.Claims{}
	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if !tkn.Valid {
		return nil, common.ErrTokenInvalid{}
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, common.ErrTokenSignatureInvalid{}
		}
		return nil, err
	}
	return claims, nil

}