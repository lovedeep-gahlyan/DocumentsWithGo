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
