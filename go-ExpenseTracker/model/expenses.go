package model

import (
	"time"
)

type Expense struct {
	ID          int       `json:"id,omitempty"`
	Description string    `json:"description"`
	Amount      int       `json:"amount"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
