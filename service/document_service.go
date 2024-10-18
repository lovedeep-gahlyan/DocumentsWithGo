package service

import (
	"docs/models"
	"docs/repositories"
	"log"
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

func (s *DocumentService) EditDocument(docID, userID uint, name, content string) *models.ResponseError {
	doc, err := s.docRepo.GetDocumentByID(docID)
	if err != nil {
		return models.NewResponseError("Document not found", 404)
	}

	if doc.OwnerID != userID {
		return models.NewResponseError("Unauthorized to edit this document", 403)
	}

	doc.Name = name
	doc.Content = content
	err = s.docRepo.UpdateDocument(doc)
	if err != nil {
		return models.NewResponseError("Error updating document", 500)
	}

	return nil
}

func (s *DocumentService) DeleteDocument(docID, userID uint) *models.ResponseError {
	doc, err := s.docRepo.GetDocumentByID(docID)
	if err != nil {
		log.Printf("Error finding document: %v", err)
		return models.NewResponseError("Document not found", 404)
	}

	if doc.OwnerID != userID {
		return models.NewResponseError("Unauthorized to delete this document", 403)
	}

	err = s.docRepo.DeleteDocument(doc)
	if err != nil {
		log.Printf("Error deleting document: %v", err)
		return models.NewResponseError("Error deleting document", 500)
	}

	return nil
}
