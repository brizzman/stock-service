package services

import (
	"context"
	"stock-service/internal/domain"
	"github.com/google/uuid"
)

type ReservationService interface {
    CreateReservation(ctx context.Context, item_id uuid.UUID, qty int) (*domain.Reservation, error)
    CommitReservation(ctx context.Context, id uuid.UUID) error
    CancelReservation(ctx context.Context, id uuid.UUID) error
    GetReservation(ctx context.Context, id uuid.UUID) (*domain.Reservation, error)
}
