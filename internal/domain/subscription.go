package domain

type Subscription struct {
	ID         int64
	UserID     int64
	Title      string
	Amount     float64
	Category   Category
	BillingDay int
}

func (s *Subscription) Validate() error {
	if s.Title == "" {
		return ErrInvalidTitle
	}
	if s.Amount <= 0 {
		return ErrInvalidAmount
	}
	if s.BillingDay < 1 || s.BillingDay > 31 {
		return ErrInvalidBillingDay
	}
	if s.UserID <= 0 {
		return ErrInvalidUserID
	}
	return nil
}
