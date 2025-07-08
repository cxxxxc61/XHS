package repository

import (
	"context"
	"github.com/cxxxxc61/study/webook/domain"
	"github.com/cxxxxc61/study/webook/repository/dao"
)

var (
	EmailcomfilctErr = dao.EmailcomfilctErr
	UserNotFoundErr  = dao.UserNotFoundErr
)

type UserRepository struct {
	dao *dao.UserDao
}

func NewUserRepository(dao *dao.UserDao) *UserRepository {
	return &UserRepository{
		dao: dao,
	}
}

func (r *UserRepository) FindEmail(c context.Context, email string) (domain.User, error) {
	u, err := r.dao.FindEmail(c, email)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		Id:       u.Id,
		Email:    u.Email,
		Password: u.Password,
	}, nil
}

func (r *UserRepository) Create(c context.Context, u domain.User) error {
	return r.dao.Insert(c, dao.User{
		Email:    u.Email,
		Password: u.Password,
	})
}
