package model

import uuid "github.com/google/uuid"

type Class struct {
	ID           uuid.UUID `json:"id"`
	RoomNo       int64     `json:"room_no"`
	Subjects     []Subject `json:"subject"`
	Fee          Fee       `json:"fee"`
	Teachers     []Teacher `json:"teacher"`
	ClassTeacher Teacher   `json:"class_teacher"`
}
