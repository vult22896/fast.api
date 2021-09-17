package services

import (
	"strconv"
	"time"

	"fast.bibabo.vn/models"
	"github.com/go-redis/cache/v8"
)

type UserService interface {
	FindOne(id int) models.User
}

type userService struct {
}

func GetIntanceUserService() UserService {
	return &userService{}
}

func (s *userService) FindOne(id int) models.User {
	key := "find_user:" + strconv.Itoa(id)
	var user models.User
	error := caching.Once(&cache.Item{
		Key:   key,
		Value: &user,
		TTL:   time.Minute * 30,
		Do: func(i *cache.Item) (interface{}, error) {
			db.Where("id", id).First(&user)
			return &user, nil
		},
	})
	if error != nil {
		panic(error)
	}
	return user
}
