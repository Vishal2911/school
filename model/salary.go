package model

import "github.com/google/uuid"

type Salary struct {
	ID     uuid.UUID `json:"id"`
	Role   string    `json:"role"`
	Amount int       `json:"amount"`
}
