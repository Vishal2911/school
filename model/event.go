package model

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID        uuid.UUID
	Name      string `json:"name"`
	Fees      *Fees
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}
