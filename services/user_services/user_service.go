package services

import (
	"strconv"
	"time"

	"fast.bibabo.vn/models"
	"github.com/go-redis/cache/v8"
	"gorm.io/gorm"
)

type UserService interface {
	FindOne(id int) models.User
}

type userService struct {
	db *gorm.DB
}

func GetIntanceUserService(db *gorm.DB) UserService {
	return &userService{db: db}
}

func (s *userService) FindOne(id int) models.User {
	key := "find_user:" + strconv.Itoa(id)
	var user models.User
	error := caching.Once(&cache.Item{
		Key:   key,
		Value: &user,
		TTL:   time.Minute * 30,
		Do: func(i *cache.Item) (interface{}, error) {
			s.db.Where("id", id).First(&user)
			return &user, nil
		},
	})
	if error != nil {
		panic(error)
	}
	return user
}
