package model

import uuid "github.com/google/uuid"

type Employee struct {
	ID         uuid.UUID `json:"id"`
	Name       Name      `json:"name"`
	Role       string    `json:"role"`
	Degree     Degree    `json:"degree"`
	Address    Address   `json:"address"`
	Contact_No Number    `json:"contact_no"`
	Email      string    `json:"email"`
	Access     string    `json:"access"`
}
