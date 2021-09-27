package middlewares

import (
	"net/http"

	services "fast.bibabo.vn/services"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/cache/v8"
	"gorm.io/gorm"
)

func Auth(db *gorm.DB, cache *cache.Cache) gin.HandlerFunc {
	return func(c *gin.Context) {
		// lấy token và kiểm tra token
		authService := services.GetInstanceAuthService(db, cache)
		if isAuth := authService.Auth(c.GetHeader("Authorization")); isAuth {
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Unauthorized"})
		}
	}
}
