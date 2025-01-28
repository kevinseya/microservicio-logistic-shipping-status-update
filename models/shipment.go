package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Shipment struct {
	ShipmentID    []byte    `gorm:"type:binary(16);primaryKey"`
	OrderID       uint      `gorm:"not null" json:"orderID"`
	CarrierID     uint      `gorm:"not null" json:"carrierID"`
	State         string    `gorm:"default:'Asigned'"`
	DateAsignment time.Time `gorm:"autoCreateTime"`
}

// Migration
func (Shipment) TableName() string {
	return "shipment"
}

// Generate UUID
func (s *Shipment) BeforeCreate(tx *gorm.DB) (err error) {
	// Convert to Bin
	uuidValue := uuid.New()
	s.ShipmentID, err = uuidValue.MarshalBinary()
	return
}

// Convert Bin to String
func (s *Shipment) GetShipmentIDAsText() string {
	uuidValue, _ := uuid.FromBytes(s.ShipmentID)
	return uuidValue.String()
}
