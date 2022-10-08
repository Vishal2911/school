package model

import uuid "github.com/google/uuid"

type Employee struct {
	ID             uuid.UUID      `json:"id"`
	Name           Name           `json:"name"`
	Role           string         `json:"role"`
	Degrees        []Degree       `json:"degree"`
	Address        Address        `json:"address"`
	ContactDetails ContactDetails `json:"contact_details"`
	Email          string         `json:"email"`
	Access         string         `json:"access"`
}
