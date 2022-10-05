package model

import "github.com/google/uuid"

type Teacher struct {
	ID         uuid.UUID
	Name       *Name    `json:"name"`
	Address    *Address `json:"address"`
	Contact_No *Number  `json:"number"`
	Email      string   `json:"email"`
	Gender     string   `json:"gender"`
}
