package services

import (
	"fast.bibabo.vn/helpers"
	"fast.bibabo.vn/lib"
	"fast.bibabo.vn/mongo_models"
	userService "fast.bibabo.vn/services/user_services"
)

type PostAttachService interface {
	AttachInfo()
	attachInfoForAnswer()
	attachInfoForProduct()
	attachInfoUser()
	buildLinkCdnMedia()
}

type postAttachService struct {
	posts  []mongo_models.Post
	userId int
}

func GetInstancePostAttachService(posts []mongo_models.Post, userId int) PostAttachService {
	return &postAttachService{
		posts:  posts,
		userId: userId,
	}
}

func (s *postAttachService) AttachInfo() {
	s.attachInfoForAnswer()
	s.attachInfoForProduct()
	s.attachInfoUser()
	s.buildLinkCdnMedia()
}

func (s *postAttachService) attachInfoForAnswer() {
	for key, post := range s.posts {
		for idx := range post.Comments {
			s.posts[key].Comments[idx].CommunityPointByAnswer = 11
		}
	}
}

func (s *postAttachService) attachInfoForProduct() {

}

func (s *postAttachService) attachInfoUser() {

	userCacheService := userService.GetInstanceUserCacheService(s.userId)
	listUserFollowing := userCacheService.ListUserFollowing()
	listGroupFollow := userCacheService.ListGroupFollow()
	listTopicFollow := userCacheService.ListTopicFollow()
	listPostLike := userCacheService.ListPostLike()

	for key, post := range s.posts {
		s.posts[key].User.IsFollow = helpers.ConvertBoolToInt(helpers.SliceIntContainsValue(listUserFollowing, post.User.ID))
		s.posts[key].Group.IsFollow = helpers.ConvertBoolToInt(helpers.SliceIntContainsValue(listGroupFollow, post.Group.ID))
		for index := range post.Topics {
			s.posts[key].Topics[index].IsFollow = helpers.ConvertBoolToInt(helpers.SliceIntContainsValue(listTopicFollow, post.Topics[index].ID))
		}
		s.posts[key].IsLike = helpers.ConvertBoolToInt(helpers.SliceIntContainsValue(listPostLike, post.ID))

		for idx, comment := range post.Comments {
			s.posts[key].Comments[idx].User.IsFollow = helpers.ConvertBoolToInt(helpers.SliceIntContainsValue(listUserFollowing, comment.User.ID))
		}
	}
}

func (s *postAttachService) buildLinkCdnMedia() {

	cdn := lib.GetInstanceCdn()

	for key, post := range s.posts {
		//link avatar user
		avatar := s.posts[key].User.Avatar
		s.posts[key].User.Avatar = cdn.GetImage(avatar, lib.FOLDER_USER, lib.SIZE_LANGE, lib.SIZE_LANGE)

		//link avatar user comment
		for idx, comment := range post.Comments {
			image := comment.Image
			avatar := comment.User.Avatar
			s.posts[key].Comments[idx].Image = cdn.GetImage(image, lib.FOLDER_QUESTION, lib.SIZE_LANGE, lib.SIZE_LANGE)
			s.posts[key].Comments[idx].User.Avatar = cdn.GetImage(avatar, lib.FOLDER_USER, lib.SIZE_LANGE, lib.SIZE_LANGE)
		}

		//link image product
		for idx, product := range post.Products {
			image := product.Image
			s.posts[key].Products[idx].Image = cdn.GetImage(image, lib.FOLDER_PRODUCT, lib.SIZE_LANGE, lib.SIZE_LANGE)
		}

		//link image post
		for idx, item := range post.Images {
			image := item.Image
			s.posts[key].Images[idx].Image = cdn.GetImage(image, lib.FOLDER_QUESTION, lib.SIZE_LANGE, lib.SIZE_LANGE)
		}

	}
}
