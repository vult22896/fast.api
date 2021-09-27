package controllers

import (
	"net/http"
	"strconv"
	"time"

	"fast.bibabo.vn/models"
	authService "fast.bibabo.vn/services"
	puService "fast.bibabo.vn/services/post_services"
	uService "fast.bibabo.vn/services/user_services"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/cache/v8"
	"gorm.io/gorm"
)

type userController struct {
	db              *gorm.DB
	cache           *cache.Cache
	userService     uService.UserService
	postUserService puService.PostUserService
}

func NewUserController(db *gorm.DB, cache *cache.Cache, userService uService.UserService, postUserService puService.PostUserService) *userController {
	return &userController{
		db:              db,
		cache:           cache,
		userService:     userService,
		postUserService: postUserService,
	}
}

func (u *userController) Index(c *gin.Context) {

	var users []models.User

	err := u.cache.Once(&cache.Item{
		Key:   "myUsers",
		Value: &users,
		TTL:   time.Hour,
		Do: func(i *cache.Item) (interface{}, error) {
			u.db.Find(&users)
			return users, nil
		},
	})
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "User founded!", "users": users})
}

func (u *userController) Me(c *gin.Context) {
	authService := authService.GetInstanceAuthService(u.db, u.cache)

	userAuthId := authService.GetUserId()

	var user models.User

	err := u.cache.Once(&cache.Item{
		Key:   "get_me_info",
		Value: &user,
		TTL:   time.Minute * 5,
		Do: func(i *cache.Item) (interface{}, error) {
			u.db.Where("id", userAuthId).First(&user)
			return user, nil
		},
	})
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
}

func (u *userController) Show(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	user := u.userService.FindOne(id)
	if user.ID == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "User not found", "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
}

func (u *userController) ListPost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	page, _ := strconv.Atoi(c.Query("page"))
	if page <= 0 {
		page = 1
	}
	posts := u.postUserService.FetchPosts(page, id)
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": posts})
}
