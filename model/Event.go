package model

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID        uuid.UUID
	Name      string `json:"name"`
	Fees      *Fees
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}
