package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"fast.bibabo.vn/database"
	"fast.bibabo.vn/models"
	authService "fast.bibabo.vn/services"
	puService "fast.bibabo.vn/services/post_services"
	uService "fast.bibabo.vn/services/user_services"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/cache/v8"
)

type UserController struct {
}

var userService = uService.GetIntanceUserService()
var postUserService = puService.GetPostUserService()

func (u *UserController) Index(c *gin.Context) {
	instanceMysql := database.GetInstanceMysql()
	db := instanceMysql.Connect()

	// var post models.Post
	// var posts []models.Post
	// db.Debug().Preload("User").Limit(10).Find(&posts)

	instanceRedis := database.GetInstanceRedis()
	mycache := instanceRedis.Caching()

	var users []models.User

	err := mycache.Once(&cache.Item{
		Key:   "myUsers",
		Value: &users,
		TTL:   time.Hour,
		Do: func(i *cache.Item) (interface{}, error) {
			db.Debug().Find(&users)
			fmt.Println("query")
			return users, nil
		},
	})
	if err != nil {
		panic(err)
	}
	// instanceMongo := database.GetInstanceMongo()
	// mongo := instanceMongo.Connect().DB("bibabo").C("product")

	// product := mongo_models.Product{
	// 	Name: "SP 1",
	// }
	// var products []mongo_models.Product
	// mongo.Find(nil).All(&products)

	// pipe := mongo.Pipe([]bson.M{{"$match": bson.M{"name": "SP1"}}})

	// collection := instanceMongo.Connect().DB("bibabo").C("post_points")

	// match := bson.M{
	// 	"$match": bson.M{
	// 		"$and": []bson.M{
	// 			{"source": 1},
	// 			{"apps": 7}}}}

	// project := bson.M{"$project": bson.M{
	// 	"user_id":     1,
	// 	"question_id": 1,
	// 	"point": bson.M{
	// 		"$divide": []interface{}{"$score", 20}},
	// 	"score_c": bson.M{
	// 		"$add": []interface{}{"$question_id", 10}}}}

	// pipe := collection.Pipe([]bson.M{match, project})
	// var result []bson.M
	// iter := pipe.Iter()
	// iter.All(&result)
	// go func() {
	// 	time.Sleep(time.Second * 10)
	// }()

	c.JSON(http.StatusOK, gin.H{"message": "User founded!", "users": users})
}

func (u *UserController) Me(c *gin.Context) {
	authService := authService.GetInstanceAuthService()

	userAuthId := authService.GetUserId()

	var user models.User

	db := database.GetInstanceMysql().Connect()

	myCache := database.GetInstanceRedis().Caching()
	err := myCache.Once(&cache.Item{
		Key:   "get_me_info",
		Value: &user,
		TTL:   time.Minute * 5,
		Do: func(i *cache.Item) (interface{}, error) {
			db.Debug().Where("id", userAuthId).First(&user)
			return user, nil
		},
	})
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
}

func (u *UserController) Show(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	user := userService.FindOne(id)
	if user.ID == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "User not found", "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
}

func (u *UserController) ListPost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	page, _ := strconv.Atoi(c.Query("page"))
	if page <= 0 {
		page = 1
	}
	posts := postUserService.FetchPosts(page, id)
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": posts})
}
