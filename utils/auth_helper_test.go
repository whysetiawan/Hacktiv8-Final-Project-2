package utils_test

import (
	"final-project-2/httpserver/models"
	"testing"
)

func TestNewAuthHelper(t *testing.T) {
	userModel := models.UserModel{
		BaseModel: models.BaseModel{
			ID: 1,
		},
		Username: "Testing",
		Email:    "test@email.com",
		Password: "123qweasd",
		Age:      23,
	}

	// accessToken, refreshToken, err := utils.NewAuthHelper().GenerateToken(&userModel)

}
