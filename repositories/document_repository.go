package repositories

import (
	"docs/models"

	"gorm.io/gorm"
)

type DocumentRepository struct {
	db *gorm.DB
}

func NewDocumentRepository(db *gorm.DB) *DocumentRepository {
	return &DocumentRepository{db: db}
}

func (repo *DocumentRepository) CreateDocument(doc *models.Document) error {
	return repo.db.Create(doc).Error
}

func (repo *DocumentRepository) GetDocumentByID(id uint) (*models.Document, error) {
	var doc models.Document
	if err := repo.db.Preload("Users").First(&doc, id).Error; err != nil {
		return nil, err
	}
	return &doc, nil
}
