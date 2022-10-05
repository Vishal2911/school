package model

import uuid "github.com/google/uuid"

type Employee struct {
	ID         uuid.UUID
	Name       *Name
	Role       string `json:"role"`
	Degree     *Degree
	Address    *Address
	Contact_No *Number
	Email      string `json:"email`
	Access     string `json:"access`
}
