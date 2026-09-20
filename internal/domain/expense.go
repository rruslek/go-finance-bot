package domain

import (
	"time"
)

type Expense struct {
	ID          int64
	UserID      int64
	Amount      float64
	Description string
	Category    Category
	Date        time.Time
}

func (e *Expense) Validate() error {
	if e.Amount <= 0 {
		return ErrInvalidAmount
	}
	if e.Category == "" {
		return ErrInvalidCategory
	}
	if e.UserID <= 0 {
		return ErrInvalidUserID
	}
	return nil
}
