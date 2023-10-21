package mredis

import (
	"context"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

var pool *redis.Pool

type RedisSetting struct {
	Host            string
	Port            int
	Username        string
	Password        string
	MaxIdle         int
	MaxActive       int
	IdleTimeout     int
	MaxConnLifetime int
	Wait            bool
}

func Init(conf *RedisSetting) {
	pool = &redis.Pool{
		MaxIdle:   conf.MaxIdle,
		MaxActive: conf.MaxActive,
		Dial: func() (redis.Conn, error) {
			rc, err := redis.Dial(
				"tcp",
				fmt.Sprintf("%s:%d", conf.Host, conf.Port),
				redis.DialUsername(conf.Username),
				redis.DialPassword(conf.Password),
			)
			if err != nil {
				return nil, err
			}
			return rc, nil
		},
		IdleTimeout:     time.Duration(conf.IdleTimeout),
		MaxConnLifetime: time.Duration(conf.MaxConnLifetime),
	}
}

func Get() redis.Conn {
	return pool.Get()
}

func GetContext(ctx context.Context) redis.Conn {
	conn, _ := pool.GetContext(ctx)
	return conn
}
