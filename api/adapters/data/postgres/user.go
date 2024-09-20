package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/serdarkalayci/membership/api/adapters/data/postgres/dao"
	"github.com/serdarkalayci/membership/api/adapters/data/postgres/mappers"
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
func (ur UserRepository) CheckUser(username string) (domain.User, error) {
	var user dao.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	row, err := ur.cp.Query(ctx, `SELECT id, username, password, email FROM users WHERE username = $1`, username) 
	if err != nil {
		return domain.User{}, err
	}
	user, err = pgx.CollectExactlyOneRow(row, pgx.RowToStructByName[dao.User])
	if err != nil {
		return domain.User{}, err
	}
	return mappers.MapUserdao2User(user), nil
}

func (ur UserRepository) AddUser(u domain.User) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	u.ID = uuid.New().String()
	_, err := ur.cp.Exec(ctx, "INSERT INTO users (id, username, password, email) VALUES ($1, $2, $3, $4)", u.ID, u.Username, u.Password, u.Email) 
	if err != nil {
		return "", err
	}
	return u.ID, err
}

func (ur UserRepository) UpdateUser(u domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := ur.cp.Query(ctx, "UPDATE users username=$2, password=$3, email=$4 where id=$1", u.ID, u.Username, u.Password, u.Email) 
	if err != nil {
		return err
	}
	return err
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