package postgres

import (
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/serdarkalayci/membership/api/domain"
)

// UserRepository holds the arangodb client and database name for methods to use
type UserRepository struct {
		cp     *pgxpool.Pool
		dbName string
	}
	
	func newUserRepository(pool *pgxpool.Pool, databaseName string) UserRepository {
		return UserRepository{
			cp:     pool,
			dbName: databaseName,
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