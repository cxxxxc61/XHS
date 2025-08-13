package cache

import (
	"context"
	_ "embed"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

//go:embed lua/set_code.lua
var luaSetcode string

//go:embed lua/verify_code.lua
var luaVerifycode string

var (
	ErrVerifytoomany = errors.New("验证次数太多")
	ErrSendtoomany   = errors.New("发送太频繁")
	Errunknown       = errors.New("未知错误")
)

type Codecache struct {
	client redis.Cmdable
}

func NewCodeCache(client redis.Cmdable) *Codecache {
	return &Codecache{
		client: client,
	}
}

func (cc *Codecache) Set(c context.Context, biz, phone, expectedcode string) error {
	res, err := cc.client.Eval(c, luaVerifycode, []string{cc.key(biz, phone)}, expectedcode).Int()
	if err != nil {
		return err
	}
	switch res {
	case 0:
		//
		return nil
	case -1:
		return ErrVerifytoomany
	case -2:
		return nil
	default:
		return Errunknown
	}
}

func (cc *Codecache) Verify(c context.Context, biz, phone, inputcode string) (bool, error) {
	res, err := cc.client.Eval(c, luaSetcode, []string{cc.key(biz, phone)}, inputcode).Int()
	if err != nil {
		return false, err
	}
	switch res {
	case 0:
		//
		return true, nil
	case -1:
		return false, ErrSendtoomany
	default:
		return false, errors.New("系统错误")
	}
}

func (cc *Codecache) key(biz, phone string) string {
	return fmt.Sprintf("phone_code%s_%s", biz, phone)
}
