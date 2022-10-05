package model

import "github.com/google/uuid"

type Fees struct {
	ID     uuid.UUID
	Name   string `json:"name"`
	Type   *Type
	Amount int `json:"amount"`
}

type Type struct {
	TutionFee      int `json:"tution_fee"`
	BusFee         int `json:"bus_fee"`
	ExaminationFee int `json:"examination_fee"`
	EventFee       int `json:"event_fee"`
}
