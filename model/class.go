package model

import uuid "github.com/google/uuid"

type Class struct {
	ID           uuid.UUID
	RoomNo       int64 `json:"roomNo"`
	Subject      []Subject
	Fees         *Fees
	Teacher      []Teacher
	ClassTeacher *Teacher
}
