package services

import (
	"context"
	"stock-service/internal/models"
	"github.com/google/uuid"
)

type ReservationService interface {
    CreateReservation(ctx context.Context, item_id uuid.UUID, qty int) (*models.Reservation, error)
    CommitReservation(ctx context.Context, id uuid.UUID) error
    CancelReservation(ctx context.Context, id uuid.UUID) error
    GetReservation(ctx context.Context, id uuid.UUID) (*models.Reservation, error)
}
