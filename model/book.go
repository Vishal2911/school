package model

import (
	uuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID        uuid.UUID `json:"id" gorm:"primaryKey"`
	BookName  string    `json:"book_name"`
	Mrp       int64     `json:"mrp"`
	Publisher string    `json:"publisher"`
	Authors   []Author  `json:"authors" gorm:"many2many:book_author;" `
}

type Author struct {
	gorm.Model
	ID            uuid.UUID     `json:"id" gorm:"primaryKey"`
	Name          Name          `json:"name" gorm:"embedded;embeddedPrefix:name_"`
	Address       Address       `json:"address" gorm:"embedded;embeddedPrefix:address_"`
	ContactDetail ContactDetail `json:"contact_detail" gorm:"embedded;embeddedPrefix:contact_"`
}
