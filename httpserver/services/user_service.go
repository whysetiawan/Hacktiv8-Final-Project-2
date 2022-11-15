package services

import (
	"final-project-2/httpserver/dto"
	"final-project-2/httpserver/models"
	"final-project-2/httpserver/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(dto *dto.UpsertUserDto) (*models.UserModel, error)
	Login(dto *dto.LoginDto) (*models.UserModel, error)
	GetUsers() (*[]models.UserModel, error)
	UpdateUser(dto *dto.UpsertUserDto, user *models.UserModel) (*models.UserModel, error)
	DeleteUser(user *models.UserModel) (bool, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(r repositories.UserRepository) *userService {
	return &userService{r}
}

func (s *userService) Register(dto *dto.UpsertUserDto) (*models.UserModel, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	dto.Password = string(hashedPassword)

	user := models.UserModel{
		Username: dto.Username,
		Email:    dto.Email,
		Password: dto.Password,
		Age:      dto.Age,
	}

	_, err = s.userRepository.Register(&user)
	if err != nil {
		return &user, err
	}
	return &user, nil
}

func (s *userService) Login(dto *dto.LoginDto) (*models.UserModel, error) {
	user := models.UserModel{
		Email:    dto.Email,
		Password: dto.Password,
	}

	_, err := s.userRepository.Login(&user)
	if err != nil {
		return &user, err
	}
	return &user, nil
}

func (s *userService) UpdateUser(dto *dto.UpsertUserDto, user *models.UserModel) (*models.UserModel, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	dto.Password = string(hashedPassword)

	userModel := models.UserModel{
		BaseModel: user.BaseModel,
		Username:  dto.Username,
		Email:     dto.Email,
		Age:       dto.Age,
		Password:  dto.Password,
	}

	user, err = s.userRepository.UpdateUser(&userModel)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *userService) GetUsers() (*[]models.UserModel, error) {
	return s.userRepository.GetUsers()
}

func (s *userService) DeleteUser(user *models.UserModel) (bool, error) {
	_, err := s.userRepository.DeleteUser(user)
	if err != nil {
		return false, err
	}
	return true, nil
}
