package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/serdarkalayci/membership/api/adapters/comm/htmx/common"
)
const secretKey = "the_most_secure_secret"
const cookieName = "membershiptoken"

// Authenticate is a middleware that checks if the request is authenticated
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// t := time.Now()
		// We can obtain the session token from the requests cookies, which come with every request
		// Initialize a new instance of `Claims`
		tknStr, err := c.Cookie(cookieName)
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, redirect to login page
				c.Redirect(http.StatusTemporaryRedirect, "/loginpage")
				return
			}
			// For any other type of error, redirect to login page
			c.Redirect(http.StatusTemporaryRedirect, "/loginpage")
			return
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
			// If token is not valid, redirect to login page
			c.Redirect(http.StatusUnauthorized, "/loginpage")
			return
		}
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				// If signature is invalid, redirect to login page
				c.Redirect(http.StatusUnauthorized, "/loginpage")
				return
			}
			// For any other type of error, redirect to login page
			c.Redirect(http.StatusUnauthorized, "/loginpage")
			return
		}
		c.Set("UserID", claims.UserID)
		c.Set("Roles", claims.Roles)
		c.Next()
		return
	}
}