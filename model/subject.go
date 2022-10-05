package model

import "github.com/google/uuid"

type Subject struct {
	ID         uuid.UUID
	Name string `json:"name"`
	Book []Book `json:"book"`
}
