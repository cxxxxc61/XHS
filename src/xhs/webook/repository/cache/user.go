package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/cxxxxc61/XHS/webook/domain"
	"github.com/redis/go-redis/v9"
)

var ErrKeyNotExist = errors.New("key not exist")

type Usercache struct {
	client     redis.Cmdable
	expiration time.Duration
}

func NewUsercache(client redis.Cmdable) *Usercache {
	return &Usercache{
		client:     client,
		expiration: time.Minute * 30,
	}
}

func (cache *Usercache) Get(c context.Context, id int64) (domain.User, error) {
	key := cache.key(id)
	val, err := cache.client.Get(c, key).Bytes()
	if err != nil {
		return domain.User{}, err
	}
	var user domain.User
	err = json.Unmarshal(val, &user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (cache *Usercache) Set(c context.Context, u domain.User) error {
	val, err := json.Marshal(u)
	if err != nil {
		return err
	}
	key := cache.key(u.Id)
	return cache.client.Set(c, key, val, cache.expiration).Err()
}

func (cache *Usercache) key(id int64) string {
	return fmt.Sprintf("user_info:%d", id)
}
