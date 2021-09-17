package database

import (
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

type DBRedis interface {
	Connect() *redis.Client
	Caching() *cache.Cache
}

type dbredis struct {
	Addr     string
	Password string
	DB       int
}

var instanceRedis *dbredis
var onceRedis sync.Once

func GetInstanceRedis() DBRedis {
	onceRedis.Do(func() {
		db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

		instanceRedis = &dbredis{
			Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       db,
		}
	})
	return instanceRedis
}

func (dbRedis *dbredis) Connect() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     dbRedis.Addr,
		Password: dbRedis.Password,
		DB:       dbRedis.DB,
	})
	return rdb
}

func (dbredis *dbredis) Caching() *cache.Cache {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": ":" + os.Getenv("REDIS_PORT"),
		},
	})

	mycache := cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})
	return mycache
}
