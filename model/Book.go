package model

import uuid "github.com/google/uuid"

type Book struct {
	ID        uuid.UUID
	BookName  string `json:"bookName"`
	Mrp       int64  `json:"mrp"`
	Publisher string `json:"publisher"`
	Author    *Author
}

type Author struct {
	Name Name
}
