package postgres

import (
	"context"
	"stock-service/internal/domain"
	"stock-service/internal/ports/outbound/postgres"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type reservationRepo struct {
	db *pgxpool.Pool
}

// NewReservationRepo returns a new ReservationRepo implementation
func NewReservationRepo(db *pgxpool.Pool) postgres.ReservationRepo {
	return &reservationRepo{db: db}
}

// GetReservation fetches a reservation by ID
func (r *reservationRepo) GetReservation(ctx context.Context, id uuid.UUID) (*domain.Reservation, error) {
	const query = `
		SELECT id, warehouse_id, item_id, qty, status
		FROM reservations
		WHERE id = $1
	`

	res := &domain.Reservation{}
	err := r.db.QueryRow(ctx, query, id).Scan(
		&res.ID,
		&res.WarehouseID,
		&res.ItemID,
		&res.Qty,
		&res.Status,
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreateReservation inserts a new reservation
func (r *reservationRepo) CreateReservation(ctx context.Context, warehouseID uuid.UUID, itemID uuid.UUID, qty int) (*domain.Reservation, error) {
	const query = `
		INSERT INTO reservations (warehouse_id, item_id, qty, status)
		VALUES ($1, $2, $3, $4)
		RETURNING id, warehouse_id, item_id, qty, status
	`

	res := &domain.Reservation{}
	err := r.db.QueryRow(ctx, query, warehouseID, itemID, qty, domain.ReservationActive).Scan(
		&res.ID,
		&res.WarehouseID,
		&res.ItemID,
		&res.Qty,
		&res.Status,
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UpdateReservation updates reservation status
func (r *reservationRepo) UpdateReservation(ctx context.Context, id uuid.UUID) error {
	const query = `
		UPDATE reservations
		SET status = $2
		WHERE id = $1
	`
	_, err := r.db.Exec(ctx, query, id, domain.ReservationCommitted)
	return err
}
