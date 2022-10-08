package model

import "github.com/google/uuid"

type Teacher struct {
	ID             uuid.UUID      `json:"id"`
	Name           Name           `json:"name"`
	Address        Address        `json:"address"`
	ContactDetails ContactDetails `json:"contact_details"`
	Email          string         `json:"email"`
	Gender         string         `json:"gender"`
}
