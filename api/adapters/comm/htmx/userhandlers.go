package htmx

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/serdarkalayci/membership/api/adapters/comm/htmx/common"
	"github.com/serdarkalayci/membership/api/application"
	"github.com/serdarkalayci/membership/api/domain"
)
const secretKey = "the_most_secure_secret"
const cookieName = "membershiptoken"

func (ws WebServer) GetUserForm(c *gin.Context) {
	c.HTML(200, "userform.html", gin.H{
		"title": "User Page",
	})
}

func (ws WebServer) UpsertUser(c *gin.Context) {
	us := application.NewUserService(ws.dbContext)
	id := c.PostForm("id")
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")
	user := domain.User{
		ID: id,
		Username: username,
		Password: password,
		Email: email,
	}
	_, err := us.UpsertUser(user)
	if err != nil {
		c.HTML(500, "messagedisplay.html", gin.H{
			"message": err.Error(),
		})
		return
	}
	c.Redirect(302, "/user")
}

func (ws WebServer) GetLoginPage(c *gin.Context) {
	c.HTML(200, "loginpage.html", gin.H{
		"title": "Member Page",
	})
}

func (ws WebServer) Login(c *gin.Context) {
	us := application.NewUserService(ws.dbContext)
	username := c.PostForm("username")
	password := c.PostForm("password")
	user, err := us.CheckUser(username, password)
	if err != nil {
		c.HTML(500, "messagedisplay.html", gin.H{
			"message": "Error while checking user",
		})
		return	
	}	
	if (user == domain.User{}) {
		c.HTML(401, "messagedisplay.html", gin.H{
			"message": "Invalid user",
		})
		return
	}
		// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(6000 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &common.Claims{
		UserID:   user.ID,
		Roles: []string{},
		UserName: user.Username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	var jwtKey = []byte(secretKey)
	// Create the JWT key used to create the signature
	tokenString, err := token.SignedString(jwtKey)
	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie(cookieName, tokenString, 3600, "/", strings.Split(c.Request.Host, ":")[0], false, true)
	c.Redirect(http.StatusFound, "/")
}

func checkLogin(c *gin.Context) (status bool, httpStatusCode int, claims *common.Claims) {
	// We can obtain the session token from the requests cookies, which come with every request
	// Initialize a new instance of `Claims`
	tknStr, err := c.Cookie(cookieName)
	status = false
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			httpStatusCode = http.StatusUnauthorized
			return
		}
		// For any other type of error, return a bad request status
		httpStatusCode = http.StatusBadRequest
		return
	}

	// Get the JWT string from the cookie
	var jwtKey = []byte(secretKey)
	// Initialize a new instance of `Claims`
	claims = &common.Claims{}
	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {

		return jwtKey, nil
	})
	if !tkn.Valid {
		httpStatusCode = http.StatusUnauthorized
		return
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			httpStatusCode = http.StatusUnauthorized
			return
		}
		httpStatusCode = http.StatusBadRequest
		return
	}
	status = true
	return
}

// Refresh swagger:route POST PUT /login Refresh
//
// Handler to refresh a JWT Token
//
// Responses:
//        200: OK
//		  400: Bad Request
//		  500: Internal Server Error
func (ws WebServer) Refresh(c *gin.Context) {
	status, _, claims := checkLogin(c)
	if status {
		// We ensure that a new token is not issued until enough time has elapsed
		// In this case, a new token will only be issued if the old token is within
		// 30 seconds of expiry. Otherwise, return a bad request status
		if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
			c.Redirect(302, "/loginpage")
			return
		}

		// Now, create a new token for the current use, with a renewed expiration time
		expirationTime := time.Now().Add(60 * time.Minute)
		claims.ExpiresAt = expirationTime.Unix()
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		var jwtKey = []byte(secretKey)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			c.Redirect(500, "/loginpage")
			return
		}

		// Set the new token as the users `token` cookie
		c.SetCookie("membershiplogin", tokenString, 3600, "/", strings.Split(c.Request.Host, ":")[0], false, true)
	} else {
		c.Redirect(401, "/loginpage")
	}
}

func (ws WebServer) Logout(c *gin.Context) {
	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie(cookieName, "", -1, "/", strings.Split(c.Request.Host, ":")[0], false, true)
	c.Redirect(http.StatusFound, "/loginpage")
}