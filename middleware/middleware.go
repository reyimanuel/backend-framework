package middleware

import (
	"backend/pkg/token"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func MiddlewareLogin(ctx *gin.Context) {
	bearerToken := ctx.GetHeader("Authorization")
	parts := strings.Split(bearerToken, " ")

	if len(parts) != 2 || parts[0] != "Bearer" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		ctx.Abort()
		return
	}
	tokenStr := parts[1]

	// Call ValidateAccessToken function from the 'token' package
	user, err := token.ValidateAccessToken(tokenStr)
	if err != nil {
		fmt.Println("Error validating token:", err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		ctx.Abort()
		return
	}

	ctx.Set("users", user)
	ctx.Next()
}
