//go:build wireinject

package wire

import (
	"github.com/cxxxxc61/XHS/wire/repository"
	"github.com/cxxxxc61/XHS/wire/repository/dao"
	"github.com/google/wire"
)

func InitRepository() *repository.UserRepository {
	wire.Build(repository.NewUserRepository, dao.NewUserDAO, Initdb)
	return new(repository.UserRepository)
}
