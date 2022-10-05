package model

import uuid "github.com/google/uuid"

type student struct {
	ID         uuid.UUID
	Name       *Name
	rollNumber int `json:"rollNumber`
	Address    *Address
	Contact_No *Number
	Email      string   `json:"email"`
	Gender     string   `json:"gender"`
	Section    string   `json:"section"`
	Parents    *Parents 
}
