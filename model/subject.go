package model

import "github.com/google/uuid"

type Subject struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Books []Book    `json:"books"`
}
