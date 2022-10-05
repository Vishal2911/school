package model

import "github.com/google/uuid"

type Facilities struct {
	ID     uuid.UUID
	Name   string `json:"name"`
	FType  string `json:"ftype`
	Amount int    `json:"amount"`
}
