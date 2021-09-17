package services

import (
	"strconv"
	"time"

	"fast.bibabo.vn/database"
	"fast.bibabo.vn/mongo_models"
	"github.com/go-redis/cache/v8"
	"gopkg.in/mgo.v2/bson"
)

type PostUserService interface {
	FetchPosts(page int, userId int) []mongo_models.Post
}

type postUserService struct {
	limit int
}

func GetPostUserService() PostUserService {
	return &postUserService{
		limit: 20,
	}
}

var mongo = database.GetInstanceMongo().Connect().DB("bibabo")
var caching = database.GetInstanceRedis().Caching()

func (s *postUserService) FetchPosts(page int, userId int) []mongo_models.Post {
	var posts []mongo_models.Post
	key := "fetch_posts:" + strconv.Itoa(userId) + "page:" + strconv.Itoa(page)
	error := caching.Once(&cache.Item{
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
	postAttachService := GetInstancePostAttachService(posts, userId)
	postAttachService.AttachInfo()

	return posts
}
