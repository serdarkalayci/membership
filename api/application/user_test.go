package application

import (
	"testing"

	apperr "github.com/serdarkalayci/membership/api/application/errors"
	"github.com/serdarkalayci/membership/api/domain"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

type mockUserRepository struct{}

var (
	getUserFunc               func(ID string) (domain.User, error)
	checkUserFunc             func(username string) (domain.User, error)
	addUserFunc               func(u domain.User) (string, error)
	updateUserFunc			  func(u domain.User) error
	addConfirmationCodeFunc   func(userID string, confirmationCode string) error
	checkConfirmationCodeFunc func(userID string, confirmationCode string) error
	activateUserFunc          func(userID string) error
)

// GetUser gets the user with the given ID
func (m mockUserRepository) GetUser(ID string) (domain.User, error) {
	return getUserFunc(ID)
}

// CheckUser checks if the user with the given email exists
func (m mockUserRepository) CheckUser(username string) (domain.User, error) {
	return checkUserFunc(username)
}

// AddUser adds a new user to the database
func (m mockUserRepository) AddUser(u domain.User) (string, error) {
	return addUserFunc(u)
}

// UpdateUser updates the user with the given ID
func (m mockUserRepository) UpdateUser(u domain.User) error {
	return updateUserFunc(u)
}

// AddConfirmationCode adds a confirmation code to the user with the given ID
func (m mockUserRepository) AddConfirmationCode(userID string, confirmationCode string) error {
	return addConfirmationCodeFunc(userID, confirmationCode)
}

// CheckConfirmationCode checks if the confirmation code is correct for the user with the given ID
func (m mockUserRepository) CheckConfirmationCode(userID string, confirmationCode string) error {
	return checkConfirmationCodeFunc(userID, confirmationCode)
}

// ActivateUser activates the user with the given ID
func (m mockUserRepository) ActivateUser(userID string) error {
	return activateUserFunc(userID)
}

func TestGetUser(t *testing.T) {
	mc := &MockContext{}
	mc.SetRepositories(&mockUserRepository{}, nil, nil)
	us := NewUserService(mc)
	getUserFunc = func(ID string) (domain.User, error) {
		return domain.User{}, apperr.ErrUserNotFound{}
	}
	_, err := us.GetUser("1")
	assert.ErrorAs(t, err, &apperr.ErrUserNotFound{})
}

func TestCheckUser(t *testing.T) {
	mc := &MockContext{}
	mc.SetRepositories(&mockUserRepository{}, nil, nil)
	us := NewUserService(mc)
	checkUserFunc = func(username string) (domain.User, error) {
		return domain.User{}, apperr.ErrUserNotFound{}
	}
	_, err := us.CheckUser("username", "password")
	assert.ErrorAs(t, err, &apperr.ErrUserNotFound{})
	hashedpass := hashPassword("password")
	checkUserFunc = func(username string) (domain.User, error) {
		return domain.User{
			ID:       username,
			Email:    "user1@email.com",
			Password: hashedpass,
		}, nil
	}
	user, err := us.CheckUser("username", "password")
	assert.NoError(t, err)
	assert.Equal(t, "user1", user.ID)
	// Check wrong password
	_, err = us.CheckUser("username", "wrongpassword")
	assert.ErrorAs(t, err, &apperr.ErrWrongCredentials{})
}

func TestCheckPassword(t *testing.T) {
	err := checkPassword("less")
	assert.ErrorAs(t, err, &apperr.ErrPasswordNotStrong{})
	assert.EqualError(t, err, "password is not strong enough. password must be at least 6 characters")
	err = checkPassword("onlylowercase")
	assert.ErrorAs(t, err, &apperr.ErrPasswordNotStrong{})
	err = checkPassword("ONLYUPPERCASE")
	assert.ErrorAs(t, err, &apperr.ErrPasswordNotStrong{})
	err = checkPassword("123456789")
	assert.ErrorAs(t, err, &apperr.ErrPasswordNotStrong{})
	err = checkPassword("lowerand123")
	assert.ErrorAs(t, err, &apperr.ErrPasswordNotStrong{})
	err = checkPassword("UPPERAND123")
	assert.ErrorAs(t, err, &apperr.ErrPasswordNotStrong{})
	err = checkPassword("lowerandUPPER")
	assert.ErrorAs(t, err, &apperr.ErrPasswordNotStrong{})
	err = checkPassword("lowerandUPPER123")
	assert.ErrorAs(t, err, &apperr.ErrPasswordNotStrong{})
	err = checkPassword("lowerandUPPER123!")
	assert.NoError(t, err)
}

func TestRandomString(t *testing.T) {
	str1 := randomString(7)
	assert.Len(t, str1, 7)
	str2 := randomString(7)
	assert.Len(t, str2, 7)
	assert.NotEqual(t, str1, str2)
}

func TestAddConfirmationCode(t *testing.T) {
	mc := &MockContext{}
	mc.SetRepositories(&mockUserRepository{}, nil, nil)
	us := NewUserService(mc)
	addConfirmationCodeFunc = func(userID string, confirmationCode string) error {
		return apperr.ErrUserNotFound{}
	}
	err := us.addConfirmationCode(domain.User{ID: "1"})
	assert.ErrorAs(t, err, &apperr.ErrUserNotFound{})
	addConfirmationCodeFunc = func(userID string, confirmationCode string) error {
		return nil
	}
	err = us.addConfirmationCode(domain.User{ID: "1"})
	assert.NoError(t, err)
}

func TestUpsertUser(t *testing.T) {
	mc := &MockContext{}
	mc.SetRepositories(&mockUserRepository{}, nil, nil)
	us := NewUserService(mc)
	// First let's try to add a user with a weak password
	addUserFunc = func(u domain.User) (string, error) {
		return "", apperr.DuplicateKeyError{}
	}
	user := domain.User{
		Email:    "test@t.t",
		Password: "P",
		FirstName:     "Test",
	}
	_, err := us.UpsertUser(user)
	assert.ErrorAs(t, err, &apperr.ErrPasswordNotStrong{})
	// Now let's fix the password, but this time UpsertUser will return duplicate key error
	user.Password = "P@ssw0rd123"
	_, err = us.UpsertUser(user)
	assert.ErrorAs(t, err, &apperr.DuplicateKeyError{})
	// Now let's fix the duplicate key error, bur adding confirmation code will return user not found error
	addUserFunc = func(u domain.User) (string, error) {
		return "1", nil
	}
	addConfirmationCodeFunc = func(userID string, confirmationCode string) error {
		return apperr.ErrUserNotFound{}
	}
	_, err = us.UpsertUser(user)
	assert.ErrorAs(t, err, &apperr.ErrUserNotFound{})
	// Now let's fix the user not found error, and everything should work
	addConfirmationCodeFunc = func(userID string, confirmationCode string) error {
		return nil
	}
	_, err = us.UpsertUser(user)
	assert.NoError(t, err)
}

func TestCheckConfirmationCode(t *testing.T) {
	mc := &MockContext{}
	mc.SetRepositories(&mockUserRepository{}, nil, nil)
	us := NewUserService(mc)
	// First let's try to check confirmation code with a user that doesn't exist
	checkConfirmationCodeFunc = func(userID string, confirmationCode string) error {
		return apperr.ConfirmationCodeError{}
	}
	err := us.CheckConfirmationCode("1", "code")
	assert.ErrorAs(t, err, &apperr.ConfirmationCodeError{})
	// Fix that error, but this time activating the user will return user not found error
	checkConfirmationCodeFunc = func(userID string, confirmationCode string) error {
		return nil
	}
	activateUserFunc = func(userID string) error {
		return apperr.ErrUserNotFound{}
	}
	err = us.CheckConfirmationCode("1", "code")
	assert.ErrorAs(t, err, &apperr.ErrUserNotFound{})
	// Fix that error, and everything should work
	activateUserFunc = func(userID string) error {
		return nil
	}
	err = us.CheckConfirmationCode("1", "code")
	assert.NoError(t, err)
}
func TestHashPassword(t *testing.T) {
	password := "Admin123!"
	hashedPassword := hashPassword(password)
	assert.NotEqual(t, password, hashedPassword)
	assert.NotEmpty(t, hashedPassword)

	// Verify that the hashed password can be compared correctly
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	assert.NoError(t, err)

	// Verify that a wrong password does not match
	wrongPassword := "WrongP@ssw0rd"
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(wrongPassword))
	assert.Error(t, err)
}

