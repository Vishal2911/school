package model

import uuid "github.com/google/uuid"

type Book struct {
	ID        uuid.UUID `json:"id"`
	BookName  string    `json:"book_name"`
	Mrp       int64     `json:"mrp"`
	Publisher string    `json:"publisher"`
	Author    Author    `json:"author"`
}

type Author struct {
	Name Name `json:"name"`
}
