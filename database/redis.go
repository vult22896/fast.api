package database

import (
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
		instanceRedis = &dbredis{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}
	})
	return instanceRedis
}

func (dbRedis *dbredis) Connect() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}

func (dbredis *dbredis) Caching() *cache.Cache {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": ":6379",
		},
	})

	mycache := cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})
	return mycache
}
