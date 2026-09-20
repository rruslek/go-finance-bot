package domain

import (
	"errors"
)

var ErrInvalidAmount = errors.New("amount must be positive")
var ErrInvalidCategory = errors.New("category is required")
var ErrInvalidTitle = errors.New("title is required")
var ErrInvalidBillingDay = errors.New("billing day must be between 1 and 31")
var ErrInvalidUserID = errors.New("user id is required")
