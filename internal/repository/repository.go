package repository

import (
	"context"
	"go-finance-bot/internal/domain"
)

type ExpenseRepository interface {
	Save(ctx context.Context, expense *domain.Expense) error
	GetByID(ctx context.Context, id int64) (*domain.Expense, error)
	ListByUserID(ctx context.Context, userID int64) ([]*domain.Expense, error)
	Delete(ctx context.Context, id int64) error
}

type SubscriptionRepository interface {
	Save(ctx context.Context, sub *domain.Subscription) error
	ListByUserID(ctx context.Context, userID int64) ([]*domain.Subscription, error)
	Delete(ctx context.Context, id int64) error
}