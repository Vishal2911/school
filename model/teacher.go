package model

import "github.com/google/uuid"

type Teacher struct {
	ID            uuid.UUID     `json:"id"`
	Name          Name          `json:"name"`
	Address       Address       `json:"address"`
	ContactDetail ContactDetail `json:"contact_detail"`
	Email         string        `json:"email"`
	Gender        string        `json:"gender"`
}
