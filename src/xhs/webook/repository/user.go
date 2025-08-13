package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/cxxxxc61/XHS/webook/domain"
	cache2 "github.com/cxxxxc61/XHS/webook/repository/cache"
	"github.com/cxxxxc61/XHS/webook/repository/dao"
)

var (
	ComfilctErr     = dao.ComfilctErr
	UserNotFoundErr = dao.UserNotFoundErr
)

type UserRepository struct {
	dao   *dao.UserDao
	cache *cache2.Usercache
}

func NewUserRepository(dao *dao.UserDao, cache *cache2.Usercache) *UserRepository {
	return &UserRepository{
		dao:   dao,
		cache: cache,
	}
}

func (r *UserRepository) Create(c context.Context, u domain.User) error {
	return r.dao.Insert(c, r.domainToentity(u))
}

func (r *UserRepository) FindEmail(c context.Context, email string) (domain.User, error) {
	u, err := r.dao.FindEmail(c, email)
	if err != nil {
		return domain.User{}, err
	}
	return r.entityTodomain(u), nil
}

func (r *UserRepository) FindId(c context.Context, id int64) (domain.User, error) {
	u, err := r.cache.Get(c, id)
	if err == nil {
		return u, nil
	}
	//if err == cache2.ErrKeyNotExist {
	//
	//}
	u1, err := r.dao.FindId(c, id)
	if err != nil {
		return domain.User{}, err
	}
	u = r.entityTodomain(u1)
	err = r.cache.Set(c, u)
	if err != nil {
		//做监控
	}
	return u, err
}

func (r *UserRepository) FindPhone(c context.Context, phone string) (domain.User, error) {
	u, err := r.dao.FindPhone(c, phone)
	if err != nil {
		return domain.User{}, err
	}
	return r.entityTodomain(u), nil
}

func (r *UserRepository) entityTodomain(u dao.User) domain.User {
	return domain.User{
		Id:       u.Id,
		Email:    u.Email.String,
		Password: u.Password,
		Phone:    u.Phone.String,
		Ctime:    time.UnixMilli(u.Ctime),
	}
}

func (r *UserRepository) domainToentity(u domain.User) dao.User {
	return dao.User{
		Id: u.Id,
		Email: sql.NullString{
			String: u.Email,
			Valid:  u.Email != "",
		},
		Password: u.Password,
		Phone: sql.NullString{
			String: u.Phone,
			Valid:  u.Phone != "",
		},
		Ctime: u.Ctime.UnixMilli(),
	}
}
