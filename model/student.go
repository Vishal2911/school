package model

import uuid "github.com/google/uuid"

type Student struct {
	ID            uuid.UUID     `json:"id"`
	Name          Name          `json:"name"`
	RollNumber    int           `json:"roll_number"`
	Address       Address       `json:"address"`
	ContactDetail ContactDetail `json:"contact_detail"`
	Email         string        `json:"email"`
	Gender        string        `json:"gender"`
	Section       string        `json:"section"`
	Parents       Parents       `json:"parents"`
}
