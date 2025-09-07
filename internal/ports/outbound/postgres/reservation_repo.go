package postgres

import (
	"context"
	"stock-service/internal/domain"
	"github.com/google/uuid"
)

type ReservationRepo interface {
	GetReservation(ctx context.Context, id uuid.UUID) (*domain.Reservation, error)
    CreateReservation(ctx context.Context, warehouseID uuid.UUID, itemID uuid.UUID, qty int) (*domain.Reservation, error)
    UpdateReservation(ctx context.Context, id uuid.UUID) error
}
