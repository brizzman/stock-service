package domain

import (
	"github.com/google/uuid"
)

type Stock struct {
	ID          uuid.UUID 
	WarehouseID uuid.UUID 
	ItemID      uuid.UUID 
	OnHand      int  
	Reserved    int
}