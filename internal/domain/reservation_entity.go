package domain

import (
	"github.com/google/uuid"
)

type ReservationStatus string

const (
	ReservationActive    ReservationStatus = "ACTIVE"
	ReservationCommitted ReservationStatus = "COMMITTED"
	ReservationCancelled ReservationStatus = "CANCELLED"
	ReservationExpired   ReservationStatus = "EXPIRED"
)

type Reservation struct {
	ID          uuid.UUID       
	WarehouseID uuid.UUID        
	ItemID      uuid.UUID     
	Qty         int             
	Status      ReservationStatus 
}