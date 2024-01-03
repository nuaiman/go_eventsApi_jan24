package middlewares

import (
	"net/http"

	"example.com/events_api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized action."})
		return
	}
	userId, err := utils.VerifyJWT(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized action."})
		return
	}
	context.Set("userId", userId)
	context.Next()
}
