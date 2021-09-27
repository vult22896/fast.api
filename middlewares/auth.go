package middlewares

import (
	"net/http"

	services "fast.bibabo.vn/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Auth(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// lấy token và kiểm tra token
		authService := services.GetInstanceAuthService(db)
		if isAuth := authService.Auth(c.GetHeader("Authorization")); isAuth {
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Unauthorized"})
		}
	}
}
