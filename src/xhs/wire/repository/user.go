package repository

import (
	"github.com/cxxxxc61/XHS/wire/repository/dao"
)

type UserRepository struct {
	dao dao.UserDAO
}

func NewUserRepository(dao dao.UserDAO) *UserRepository {
	return &UserRepository{
		dao: dao,
	}
}
