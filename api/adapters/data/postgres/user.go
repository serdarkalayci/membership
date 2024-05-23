package postgres

import (
	"errors"

	"github.com/serdarkalayci/membership/api/domain"
	"gorm.io/gorm"
)

// UserRepository holds the arangodb client and database name for methods to use
type UserRepository struct {
	db *gorm.DB
}

func newUserRepository(database *gorm.DB) UserRepository {
	return UserRepository{
		db: database,
	}
}

func (ur UserRepository) GetUser(ID string) (domain.User, error) {
	return domain.User{}, errors.New("Not implemented")
}
func (ur UserRepository) CheckUser(email string) (domain.User, error) {
	return domain.User{}, errors.New("Not implemented")
}
func (ur UserRepository) AddUser(u domain.User) (string, error) {
	return "", errors.New("Not implemented")
}
func (ur UserRepository) AddConfirmationCode(userID string, confirmationCode string) error {
	return errors.New("Not implemented")
}
func (ur UserRepository) CheckConfirmationCode(userID string, confirmationCode string) error {
	return errors.New("Not implemented")
}
func (ur UserRepository) ActivateUser(userID string) error {
	return errors.New("Not implemented")
}