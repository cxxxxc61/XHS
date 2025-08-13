package service

import (
	"context"
	"errors"

	"github.com/cxxxxc61/XHS/webook/domain"
	"github.com/cxxxxc61/XHS/webook/repository"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

var (
	EmailcomfilctErr  = repository.ComfilctErr
	PasswordorUserErr = errors.New("账号/邮箱或密码不对")
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) Login(c context.Context, u domain.User) (domain.User, error) {
	uc, err := svc.repo.FindEmail(c, u.Email)
	if err == repository.UserNotFoundErr {
		return domain.User{}, PasswordorUserErr
	}
	if err != nil {
		return domain.User{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(uc.Password), []byte(u.Password))
	if err != nil {
		return domain.User{}, PasswordorUserErr
	}
	return u, nil
}

func (svc *UserService) Signup(ctx context.Context, u domain.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return svc.repo.Create(ctx, u)
}

func (svc *UserService) Profile(c context.Context, id int64) (domain.User, error) {
	u, err := svc.repo.FindId(c, id)
	if err == nil {
		return u, nil
	}
	if err == redis.Nil {

	}
	return u, nil
}

func (svc *UserService) FindorCreate(c context.Context, phone string) (domain.User, error) {
	u, err := svc.repo.FindPhone(c, phone)
	if err != repository.UserNotFoundErr {
		return u, err
	}
	err = svc.repo.Create(c, domain.User{
		Phone: phone,
	})
	if err != nil && err != repository.ComfilctErr {
		return domain.User{}, err
	}
	return svc.repo.FindPhone(c, phone)
}
