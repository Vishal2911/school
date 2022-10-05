package model

import "github.com/google/uuid"

type Fees struct {
	ID         uuid.UUID
	Name   string `json:"name"`
	Type   *Type  
	Amount int    `json:"amount"`
}

type Type struct {
  TutionFee int `json:"tutionfee"`
  BusFee int `json:"busfee"`
  ExaminationFee int `json:"examinationFee"`
  EventFee int `json:"eventFee"`

}
