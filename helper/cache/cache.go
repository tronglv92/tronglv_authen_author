package cache

import (
	"fmt"
	"github/tronglv_authen_author/helper/errors"
	"github/tronglv_authen_author/helper/stores/redis"

	"log"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/syncx"
)

const (
	RedisDriver string = "redis"
)

type Config struct {
	Stack string            `json:"stack,optional"`
	Redis redis.RedisConfig `json:"redis,optional"`
}

type Cache interface {
	cache.Cache
}

func New(c Config, opts ...Option) Cache {
	switch c.Stack {
	case RedisDriver:
		return NewNode(
			redis.MustNewRedis(c.Redis),
			syncx.NewSingleFlight(),
			cache.NewStat(RedisDriver),
			errors.InternalServer(fmt.Errorf(RedisDriver)),
			opts...,
		)
	}
	log.Fatal("no cache driver support")
	return nil
}
