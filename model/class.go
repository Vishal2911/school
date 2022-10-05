package model

import uuid "github.com/google/uuid"

type Class struct {
	ID           uuid.UUID
	RoomNo       int64 `json:"room_no"`
	Subject      []Subject
	Fees         *Fees
	Teacher      []Teacher
	ClassTeacher *Teacher
}
