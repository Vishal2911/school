package model

import "github.com/google/uuid"

type Salary struct {
	ID         uuid.UUID
	Role string `json:"role"`
	Amount int `json:"amount"`
}
