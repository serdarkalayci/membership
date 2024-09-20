package mappers

import (
	"github.com/serdarkalayci/membership/api/adapters/data/postgres/dao"
	"github.com/serdarkalayci/membership/api/domain"
)

func MapUserdao2User(userdao dao.User) domain.User {
	return domain.User{
		ID: userdao.ID.String(),
		Username: userdao.Username,
		Password: userdao.Password,
		Email: userdao.Email,
	}
}