package model

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Fees      Fee      `json:"fees"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}
