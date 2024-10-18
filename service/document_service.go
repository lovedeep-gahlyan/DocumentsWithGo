package service

import (
	"docs/models"
	"docs/repositories"
)

type DocumentService struct {
	docRepo  *repositories.DocumentRepository
	userRepo *repositories.UserRepository
}

func NewDocumentService(dr *repositories.DocumentRepository, ur *repositories.UserRepository) *DocumentService {
	return &DocumentService{docRepo: dr, userRepo: ur}
}

func (s *DocumentService) CreateDocument(userID uint, name, content string) (*models.Document, *models.ResponseError) {
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, models.NewResponseError("User not found", 404)
	}

	document := &models.Document{
		Name:    name,
		Content: content,
		OwnerID: user.ID,
		Owner:   *user,
		Users:   []models.User{*user},
	}

	err = s.docRepo.CreateDocument(document)
	if err != nil {
		return nil, models.NewResponseError("Error creating document", 500)
	}

	return document, nil
}
