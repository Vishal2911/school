package model

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Fee       Fee       `json:"fees" gorm:"embedded;embeddedPrefix:fee_"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}
