package services

import (
	"strconv"
	"time"

	"fast.bibabo.vn/database"
	"fast.bibabo.vn/mongo_models"
	"github.com/go-redis/cache/v8"
	"gopkg.in/mgo.v2/bson"
	"gorm.io/gorm"
)

type PostUserService interface {
	FetchPosts(page int, userId int) []mongo_models.Post
}

type postUserService struct {
	limit int
	db    *gorm.DB
	cache *cache.Cache
}

func GetPostUserService(db *gorm.DB, cache *cache.Cache) PostUserService {
	return &postUserService{
		limit: 20,
		db:    db,
		cache: cache,
	}
}

var mongo = database.GetInstanceMongo().Connect().DB("bibabo")

func (s *postUserService) FetchPosts(page int, userId int) []mongo_models.Post {
	var posts []mongo_models.Post
	key := "fetch_posts:" + strconv.Itoa(userId) + "page:" + strconv.Itoa(page)
	error := s.cache.Once(&cache.Item{
		Key:   key,
		Value: &posts,
		TTL:   time.Minute * 5,
		Do: func(i *cache.Item) (interface{}, error) {
			collection := mongo.C("posts")
			collection.Find(bson.M{"user.id": userId}).Skip((page - 1) * s.limit).Limit(s.limit).All(&posts)
			return &posts, nil
		},
	})
	if error != nil {
		panic(error)
	}
	postAttachService := GetInstancePostAttachService(posts, userId, s.db, s.cache)
	postAttachService.AttachInfo()

	return posts
}
