package service

import (
	"docs/models"
	"docs/repositories"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(username string) (*models.User, *models.ResponseError) {
	user := &models.User{
		Username: username,
	}

	err := s.userRepo.CreateUser(user)
	if err != nil {
		return nil, models.NewResponseError("Error creating user", 500)
	}

	return user, nil
}
