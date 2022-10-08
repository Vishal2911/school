package model

import "github.com/google/uuid"

type Facilities struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Type   string    `json:"type"`
	Amount int       `json:"amount"`
}
