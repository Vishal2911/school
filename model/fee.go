package model

import "github.com/google/uuid"

type Fee struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Type   Type      `json:"type"`
	Amount int       `json:"amount"`
}

type Type struct {
	TutionFee      int `json:"tution_fee"`
	BusFee         int `json:"bus_fee"`
	ExaminationFee int `json:"examination_fee"`
	EventFee       int `json:"event_fee"`
}
