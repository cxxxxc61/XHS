package repository

import (
	"context"

	"github.com/cxxxxc61/XHS/webook/repository/cache"
)

var (
	ErrSendtoomany   = cache.ErrSendtoomany
	ErrVerifytoomany = cache.ErrVerifytoomany
)

type CodeRepository struct {
	cache *cache.Codecache
}

func NewCodeRepository(cc *cache.Codecache) *CodeRepository {
	return &CodeRepository{
		cache: cc,
	}
}

func (cr *CodeRepository) Store(c context.Context, biz string,
	phone string, code string) error {
	return cr.cache.Set(c, biz, phone, code)
}

func (cr *CodeRepository) Verify(c context.Context, biz string,
	phone string, inputcode string) (bool, error) {
	return cr.cache.Verify(c, biz, phone, inputcode)
}
