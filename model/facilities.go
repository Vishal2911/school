package model

import "github.com/google/uuid"

type Facilities struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	FType  string    `json:"f_type"`
	Amount int       `json:"amount"`
}
