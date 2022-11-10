package middleware

import (
	"encoding/json"
	"final-project-2/httpserver/services"
	"final-project-2/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtGuard(s services.AuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken := ctx.Request.Header.Get("Authorization")
		hasPrefix := strings.HasPrefix(accessToken, "Bearer")

		if !hasPrefix {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.NewHttpError("Unauthorized", "No Bearer Found"))
			return
		}

		splitToken := strings.Split(accessToken, " ")

		if len(splitToken) < 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.NewHttpError("Unauthorized", "Invalid Token"))
			return
		}

		jwtToken := splitToken[1]

		isVerified, jwtDecoded, err := s.VerifyToken(string(jwtToken))

		if !isVerified {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.NewHttpError("Unauthorized", err.Error()))
			return
		}
		a, err := json.Marshal(jwtDecoded)
		fmt.Println(string(a))
		print(err)
		ctx.Set("user", jwtDecoded)
		ctx.Next()
	}
}
