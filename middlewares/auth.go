package middlewares

import (
	"net/http"

	services "fast.bibabo.vn/services"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	// lấy token và kiểm tra token
	authService := services.GetInstanceAuthService()
	if isAuth := authService.Auth(c.GetHeader("Authorization")); isAuth {
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Unauthorized"})
	}
}
