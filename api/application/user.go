// Package application is the package that holds the application logic between database and communication layers
package application

import (
	"fmt"
	"math/rand"
	"unicode"

	"github.com/rs/zerolog/log"
	apperr "github.com/serdarkalayci/membership/api/application/errors"
	"github.com/serdarkalayci/membership/api/domain"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

// UserRepository is the interface to interact with User domain object
type UserRepository interface {
	GetUser(ID string) (domain.User, error)
	CheckUser(username string) (domain.User, error)
	AddUser(u domain.User) (string, error)
	UpdateUser(u domain.User) error
	AddConfirmationCode(userID string, confirmationCode string) error
	CheckConfirmationCode(userID string, confirmationCode string) error
	ActivateUser(userID string) error
}

// UserService is the struct to let outer layers to interact to the User Applicatopn
type UserService struct {
	dc DataContextCarrier
}

// NewUserService creates a new UserService instance and sets its repository
func NewUserService(dc DataContextCarrier) UserService {
	return UserService{
		dc: dc,
	}
}

// GetUser simply returns a single user or an error that is returned from the repository
func (us UserService) GetUser(ID string) (domain.User, error) {
	return us.dc.GetUserRepository().GetUser(ID)
}

// CheckUser checks if the username and password maches any from the repository by first hashing its password, returns error if none found
func (us UserService) CheckUser(username, password string) (domain.User, error) {
	log.Info().Msgf("CheckUser: %v", username)
	user, err := us.dc.GetUserRepository().CheckUser(username)
	if err != nil {
		log.Warn().Msgf("Error while checking user: %v", err)
		return domain.User{}, err
	}
	if comparePasswords(password, user.Password) {
		log.Info().Msgf("User found: %v", user)
		return user, nil
	}
	log.Warn().Msgf("Error while checking user: %v", apperr.ErrWrongCredentials{})
	return domain.User{}, apperr.ErrWrongCredentials{}
}

// UpsertUser inserts or updates a user to the repository by first hashing its password
func (us UserService) UpsertUser(u domain.User) (string, error) {
	if err := checkPassword(u.Password); err != nil {
		return "", err
	}
	u.Password = hashPassword(u.Password)
	if u.ID != "" { // Update
		err := us.dc.GetUserRepository().UpdateUser(u)
		if err != nil {
			return "", err
		}
		return u.ID, nil
	}
	// Add
	newUID, err := us.dc.GetUserRepository().AddUser(u)
	if err != nil {
		return "", err
	}
	// // Generate a random string and send an email to the user with the confirmation code
	// u.ID = newUID
	// err = us.addConfirmationCode(u)
	// if err != nil {
	// 	return err
	// }
	return newUID, nil
}

// CheckConfirmationCode checks if the confirmation code matches the one in the repository, if so, activates the user
func (us UserService) CheckConfirmationCode(userID string, confirmationCode string) error {
	err := us.dc.GetUserRepository().CheckConfirmationCode(userID, confirmationCode)
	if err != nil {
		return err
	}
	err = us.dc.GetUserRepository().ActivateUser(userID)
	if err != nil {
		return err
	}
	return nil
}

func (us UserService) addConfirmationCode(u domain.User) error {
	confirmationCode := randomString(7)
	err := us.dc.GetUserRepository().AddConfirmationCode(u.ID, confirmationCode)
	if err != nil {
		return err
	}
	// Send email to user with the confirmation code
	sendConfirmationEmail(u, confirmationCode)
	return nil
}

func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func sendConfirmationEmail(u domain.User, confirmationCode string) error {
	to := u.Email
	subject := viper.GetViper().GetString("ConformationCodeSubject")
	body := fmt.Sprintf(viper.GetString("ConfirmationCodeMessage"), u.FirstName, confirmationCode, u.ID)
	return sendEmail(to, subject, body)
}

// HashPassword hashes the password string in order to get ready to store or check if it matches the stored value
func hashPassword(password string) string {
	// Convert the password to a byte slice
	passwordBytes := []byte(password)
	// Generate the bcrypt hash of the password
	hash, _ := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	// Convert the hash to a string and return it
	hashString := string(hash)
	return hashString
}

// comparePasswords compares the plaintext password with the hashed password
func comparePasswords(plaintextPassword string, hashedPassword string) bool {
	// Convert the hashed password from string to byte slice
	hashedPasswordBytes := []byte(hashedPassword)

	// Compare the plaintext password with the hashed password
	err := bcrypt.CompareHashAndPassword(hashedPasswordBytes, []byte(plaintextPassword))
	if err != nil {
		return false
	}

	return true
}

func checkPassword(password string) error {
	if len(password) < 6 {
		return apperr.ErrPasswordNotStrong{Reason: "şifre en az 6 karakterden oluşmalıdır."}
	}
next:
	for name, classes := range map[string][]*unicode.RangeTable{
		"büyük harf": {unicode.Upper, unicode.Title},
		"küçük harf": {unicode.Lower},
		"sayısal":    {unicode.Number, unicode.Digit},
		"özel":    {unicode.Space, unicode.Symbol, unicode.Punct, unicode.Mark},
	} {
		for _, r := range password {
			if unicode.IsOneOf(classes, r) {
				continue next
			}
		}
		return apperr.ErrPasswordNotStrong{Reason: fmt.Sprintf("şifrede en az bir %s karakter olmalıdır.", name)}
	}
	return nil
}
