package models

type Document struct {
	ID      uint `gorm:"primaryKey"`
	Name    string
	Content string
	OwnerID uint
	Owner   User   `gorm:"foreignKey:OwnerID"`
	Users   []User `gorm:"many2many:document_users"`
}
