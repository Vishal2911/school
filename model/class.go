package model

import uuid "github.com/google/uuid"

type Class struct {
	ID           uuid.UUID `json:"id"`
	RoomNo       int64     `json:"room_no"`
	Subject      []Subject `json:"subject"`
	Fees         Fee       `json:"fees"`
	Teacher      []Teacher `json:"teacher"`
	ClassTeacher Teacher   `json:"class_teacher"`
}
