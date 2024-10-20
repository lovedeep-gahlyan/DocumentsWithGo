package models

type Document struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	Content string `json:"content"`
	OwnerID uint   `json:"owner_id"`
	Owner   User   `gorm:"foreignKey:OwnerID;constraint:OnDelete:CASCADE;"`
	Users   []User `gorm:"many2many:document_users;constraint:OnDelete:CASCADE;"`
}
