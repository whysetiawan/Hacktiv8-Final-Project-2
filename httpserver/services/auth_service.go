package services

import (
	"encoding/json"
	"errors"
	"final-project-2/httpserver/models"
	"final-project-2/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type AuthService interface {
	VerifyToken(c *gin.Context) (interface{}, error)
	GenerateToken(user *models.UserModel) (string, string, error)
}

type authService struct {
	JWT_SECRET_KEY string
}

func NewAuthService() *authService {
	return &authService{
		JWT_SECRET_KEY: utils.Constants.JWT_SECRET_KEY,
	}
}

func (s *authService) VerifyToken(c *gin.Context) (interface{}, error) {
	errResponse := errors.New("Unauthorized")
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errResponse
	}

	if len(strings.Split(headerToken, " ")) < 2 {
		return nil, errResponse
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, err := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}

		return []byte(s.JWT_SECRET_KEY), nil
	})

	if token == nil {
		return nil, errResponse
	}

	if err != nil {
		v, _ := err.(*jwt.ValidationError)
		if v.Errors == jwt.ValidationErrorExpired {
			return nil, errResponse
		}
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResponse
	}

	return token.Claims.(jwt.MapClaims), nil
}

func (s *authService) GenerateToken(user *models.UserModel) (string, string, error) {
	const ttlAccessToken = 24 * time.Hour
	const ttlRefreshToken = (24 * 7) * time.Hour

	var userMap map[string]interface{}
	data, err := json.Marshal(user)
	if err != nil {
		return "", "", err
	}

	json.Unmarshal(data, &userMap)

	accessClaims, refreshClaims := jwt.MapClaims{
		"data": userMap,
		"exp":  time.Now().UTC().Add(ttlAccessToken).Unix(),
	}, jwt.MapClaims{
		"data": userMap,
		"exp":  time.Now().UTC().Add(ttlRefreshToken).Unix(),
	}

	var secretKeyByte = []byte(s.JWT_SECRET_KEY)

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString(secretKeyByte)

	if err != nil {
		return "", "", err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(secretKeyByte)

	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
