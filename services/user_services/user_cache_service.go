package services

import (
	"strconv"
	"time"

	"fast.bibabo.vn/models"
	"github.com/go-redis/cache/v8"
)

type UserCacheService interface {
	ListUserFollowing() []int
	ListGroupFollow() []int
	ListTopicFollow() []int
	ListPostLike() []int
	ListPostSave() []int
}

type userCacheService struct {
	userId int
}

func GetInstanceUserCacheService(userId int) UserCacheService {
	return &userCacheService{userId: userId}
}

func (s *userCacheService) ListUserFollowing() []int {
	var userIds []int
	var ufriends []models.Ufriend
	key := "list_user_following:userId-" + strconv.Itoa(s.userId)
	error := caching.Once(&cache.Item{
		Key:   key,
		Value: &userIds,
		TTL:   time.Minute * 5,
		Do: func(i *cache.Item) (interface{}, error) {
			db.Where("user_id", s.userId).Find(&ufriends)
			for _, u := range ufriends {
				userIds = append(userIds, u.UfriendFriendId)
			}

			return &userIds, nil
		},
	})
	if error != nil {
		panic(error)
	}
	return userIds
}

func (s *userCacheService) ListGroupFollow() []int {
	var groupIds []int
	var userGroupFollows []models.UserGroupFollow
	key := "list_group_follow:user_id-" + strconv.Itoa(s.userId)
	error := caching.Once(&cache.Item{
		Key:   key,
		Value: &groupIds,
		TTL:   time.Minute * 5,
		Do: func(i *cache.Item) (interface{}, error) {
			db.Where("user_id", s.userId).Where("type", 2).Find(&userGroupFollows)
			for _, item := range userGroupFollows {
				groupIds = append(groupIds, item.GroupId)
			}
			return &groupIds, nil
		},
	})
	if error != nil {
		panic(error)
	}
	return groupIds
}

func (s *userCacheService) ListTopicFollow() []int {
	var topicIds []int
	var topicLikes []models.TopicLike
	key := "list_topic_Like:user_id-" + strconv.Itoa(s.userId)
	error := caching.Once(&cache.Item{
		Key:   key,
		Value: &topicIds,
		TTL:   time.Minute * 5,
		Do: func(i *cache.Item) (interface{}, error) {
			db.Where("user_id", s.userId).Where("qa_topic_like_is_like", 1).Find(&topicLikes)
			for _, item := range topicLikes {
				topicIds = append(topicIds, item.TopicId)
			}
			return &topicIds, nil
		},
	})
	if error != nil {
		panic(error)
	}
	return topicIds
}

func (s *userCacheService) ListPostLike() []int {
	var postIds []int
	var postLikes []models.PostLike
	key := "list_post_list:user_id-" + strconv.Itoa(s.userId)
	error := caching.Once(&cache.Item{
		Key:   key,
		Value: &postIds,
		TTL:   time.Minute * 5,
		Do: func(i *cache.Item) (interface{}, error) {
			db.Where("user_id", s.userId).Where("question_like_is_like", 1).Find(&postLikes)
			for _, item := range postLikes {
				postIds = append(postIds, item.PostId)
			}
			return &postIds, nil
		},
	})
	if error != nil {
		panic(error)
	}
	return postIds
}

func (s *userCacheService) ListPostSave() []int {
	var a []int
	return a
}
