package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Shipment struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	OrderID       uuid.UUID `gorm:"type:uuid;not null" json:"order_id"`
	UserCarrierID uuid.UUID `gorm:"type:uuid;not null" json:"user_carrier_id"`
	Status        string    `gorm:"type:string;default:'ASSIGNED CARRIER'" json:"status"`
	CreatedAt     time.Time `gorm:"type:timestamp;default:current_timestamp" json:"created_at"`
	UpdatedAt     time.Time `gorm:"type:timestamp;default:current_timestamp" json:"updated_at"`
}

// Migration
func (Shipment) TableName() string {
	return "shipments"
}

// Generate UUID
func (s *Shipment) BeforeCreate(tx *gorm.DB) (err error) {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	return
}

func (s *Shipment) BeforeUpdate(tx *gorm.DB) (err error) {
	s.UpdatedAt = time.Now()
	return nil
}
