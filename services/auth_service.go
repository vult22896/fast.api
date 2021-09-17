package services

import (
	"strings"
	"sync"
	"time"

	"fast.bibabo.vn/database"
	"fast.bibabo.vn/models"
	"github.com/go-redis/cache/v8"
)

type AuthService interface {
	GetUserId() int
	Auth(string) bool
}

type authService struct {
	userId int
}

var instanceAuthService *authService
var onceAuthService sync.Once

func GetInstanceAuthService() AuthService {
	onceAuthService.Do(func() {
		instanceAuthService = &authService{}
	})
	return instanceAuthService
}

func (s *authService) GetUserId() int {
	return s.userId
}

func (s *authService) Auth(token string) bool {
	//Support both Bearer and without Bearer
	i := strings.Index(token, "Bearer")
	if i > -1 {
		len := i + len("Bearer")
		token = token[len+1:]
	}
	if len(token) == 0 {
		return false
	}
	var user_session models.UserSession

	caching := database.GetInstanceRedis().Caching()

	key := "user_session:" + token
	error := caching.Once(&cache.Item{
		Key:   key,
		Value: &user_session,
		TTL:   time.Minute * 30,
		Do: func(i *cache.Item) (interface{}, error) {
			instanceSql := database.GetInstanceMysql()
			db := instanceSql.Connect()
			db.Where("token", token).First(&user_session)
			return &user_session, nil
		},
	})
	if error != nil {
		return false
	}
	if user_session.ID > 0 {
		s.userId = user_session.UserId
		return true
	}
	return false
}
