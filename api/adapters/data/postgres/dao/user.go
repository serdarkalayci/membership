package dao

import "github.com/google/uuid"

type User struct {
	ID uuid.UUID `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Email string `db:"email"`
}