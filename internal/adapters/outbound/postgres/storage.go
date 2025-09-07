package postgres

import (
	"context"
	"errors"
	"fmt"
	"stock-service/internal/infrastructure/logger"
	"time"
	"stock-service/internal/ports/outbound/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

// Storage contains all repositories
type Storage struct {
	db          *pgxpool.Pool
	dsn 		string
	log 		logger.Logger 

	Reservation postgres.ReservationRepo
}

// NewStorage initializes the storage with pgx pool
func NewStorage(
	ctx context.Context, 
	log logger.Logger, 
	dsn string,
) (*Storage, error) {
	// Connect to PostgreSQL using pgxpool
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	// Initialize repositories with db pool
	reservationRepo := NewReservationRepo(pool)

	return &Storage{
		db:          pool,
		dsn: 		 dsn,
		log:		 log,
		Reservation: reservationRepo,
	}, nil
}

// Close closes the db connection pool
func (s *Storage) Close() {
	if s.db != nil {
		s.db.Close()
	}
}

func (s *Storage) HealthCheck(ctx context.Context) error {
	if s.db == nil {
		return errors.New("no database pool")
	}
	return s.db.Ping(ctx)
}


// KeepAlivePollPeriod is a Pg keepalive check time period
const KeepAlivePollPeriod = 3 * time.Second

// KeepAlivePg makes sure PostgreSQL is alive and reconnects if needed
func (s *Storage) KeepAlivePg(ctx context.Context) {
	ticker := time.NewTicker(KeepAlivePollPeriod)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			s.log.Info("KeepAlivePg stopped")
			return
		case <-ticker.C:
			// Check if current pool is still alive
			if err := s.HealthCheck(ctx); err == nil {
				continue
			}	

			s.log.Warn("Postgres ping failed, reconnecting...")

			newPool, err := pgxpool.New(ctx, s.dsn)
			if err != nil {
				s.log.Warn("Failed to reconnect to Postgres", zap.Error(err))
				continue
			}

			if s.db != nil {
				s.db.Close()
			}

			s.db = newPool
			s.log.Info("Reconnected to Postgres")
		}
	}
}

