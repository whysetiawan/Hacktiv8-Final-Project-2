package middleware

import (
	"final-project-2/httpserver/services"
	"final-project-2/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtGuard(s services.AuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken := ctx.Request.Header.Get("Authorization")
		hasPrefix := strings.HasPrefix(accessToken, "Bearer")

		if !hasPrefix {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.NewHttpError("Unauthorized", nil))
			return
		}

		splitToken := strings.Split(accessToken, " ")

		if len(splitToken) < 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.NewHttpError("Unauthorized", nil))
			return
		}

		jwtToken := splitToken[1]

		isVerified, user, err := s.VerifyToken(string(jwtToken))

		if !isVerified {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.NewHttpError("Unauthorized", err.Error()))
			return
		}

		ctx.Set("user", user)
		ctx.Next()
	}
}
