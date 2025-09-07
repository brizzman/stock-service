package services

import (
	"context"
	"fmt"
	"stock-service/internal/domain"
	"stock-service/internal/application/ports/outbound/persistence"

	"github.com/google/uuid"
)

type reservationService struct {
	repo persistence.ReservationRepo
}

// NewReservationService returns a ReservationService implementation
func NewReservationService(repo persistence.ReservationRepo) ReservationService {
	return &reservationService{repo: repo}
}

func (s *reservationService) CreateReservation(ctx context.Context, itemID uuid.UUID, qty int) (*domain.Reservation, error) {
	// ⚠️ To implement FindAvailable() warehouse 
	warehouseID := uuid.New()

	res, err := s.repo.CreateReservation(ctx, warehouseID, itemID, qty)
	if err != nil {
		return nil, fmt.Errorf("failed to create reservation: %w", err)
	}
	return res, nil
}

func (s *reservationService) CommitReservation(ctx context.Context, id uuid.UUID) error {
	err := s.repo.UpdateReservation(ctx, id, domain.ReservationCommitted)
	if err != nil {
		return fmt.Errorf("failed to commit reservation %s: %w", id, err)
	}
	return nil
}

func (s *reservationService) CancelReservation(ctx context.Context, id uuid.UUID) error {
	err := s.repo.UpdateReservation(ctx, id, domain.ReservationCancelled)
	if err != nil {
		return fmt.Errorf("failed to cancel reservation %s: %w", id, err)
	}
	return nil
}

func (s *reservationService) GetReservation(ctx context.Context, id uuid.UUID) (*domain.Reservation, error) {
	res, err := s.repo.GetReservation(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get reservation %s: %w", id, err)
	}
	return res, nil
}
