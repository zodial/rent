package models

import (
	"time"

	"gorm.io/gorm"
)

// Administrator represents system users with roles.
type Administrator struct {
	ID        uint   `gorm:"primaryKey"`
	TenantID  string `gorm:"index;not null;default:'default'"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"uniqueIndex;not null"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"index;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Property is a building or complex containing rooms.
type Property struct {
	ID        uint   `gorm:"primaryKey"`
	TenantID  string `gorm:"index;not null;default:'default'"`
	Name      string `gorm:"not null"`
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Rooms     []Room `gorm:"foreignKey:PropertyID"`
}

// PropertyAdmin links administrators to properties with equity percent.
type PropertyAdmin struct {
	ID              uint   `gorm:"primaryKey"`
	TenantID        string `gorm:"index;not null;default:'default'"`
	PropertyID      uint   `gorm:"index;not null"`
	AdministratorID uint   `gorm:"index;not null"`
	EquityPercent   int64  `gorm:"not null"` // stored as basis points (10000 = 100.00%)
	CreatedAt       time.Time
}

// Room represents a rentable room within a property.
type Room struct {
	ID         uint   `gorm:"primaryKey"`
	TenantID   string `gorm:"index;not null;default:'default'"`
	PropertyID uint   `gorm:"index;not null"`
	Number     string `gorm:"not null"`
	Area       float64
	RentAmount int64  `gorm:"not null"` // cents
	Status     string `gorm:"not null;default:'vacant'"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// Tenant is the contract signee.
type Tenant struct {
	ID        uint   `gorm:"primaryKey"`
	TenantID  string `gorm:"index;not null;default:'default'"`
	Name      string `gorm:"not null"`
	Contact   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Contract represents a signed rental agreement.
type Contract struct {
	ID           uint   `gorm:"primaryKey"`
	TenantID     string `gorm:"index;not null;default:'default'"`
	RoomID       uint   `gorm:"index;not null"`
	TenantRef    uint   `gorm:"not null"` // tenant who signed
	StartDate    time.Time
	EndDate      *time.Time
	RentAmount   int64 `gorm:"not null"`
	PaymentCycle string
	Status       string `gorm:"not null;default:'draft'"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// Transaction and Allocation for settlements.
type Transaction struct {
	ID         uint   `gorm:"primaryKey"`
	TenantID   string `gorm:"index;not null;default:'default'"`
	ContractID uint   `gorm:"index;not null"`
	Amount     int64  `gorm:"not null"` // cents
	DueDate    *time.Time
	Paid       bool `gorm:"default:false"`
	CreatedAt  time.Time
}

type Allocation struct {
	ID              uint   `gorm:"primaryKey"`
	TenantID        string `gorm:"index;not null;default:'default'"`
	TransactionID   uint   `gorm:"index;not null"`
	PropertyID      uint   `gorm:"index;not null"`
	AdministratorID uint   `gorm:"index;not null"`
	Amount          int64  `gorm:"not null"`
	RoundingAdj     int64  `gorm:"not null"`
	CreatedAt       time.Time
}

// AuditLog for tracing actions.
type AuditLog struct {
	ID         uint `gorm:"primaryKey"`
	TenantID   string
	ActorID    uint
	Action     string
	TargetType string
	TargetID   uint
	Diff       string
	CreatedAt  time.Time
}

// Migrate runs schema migrations for the MVP.
func Migrate(gdb *gorm.DB) error {
	return gdb.AutoMigrate(
		&Administrator{},
		&Property{},
		&PropertyAdmin{},
		&Room{},
		&Tenant{},
		&Contract{},
		&Transaction{},
		&Allocation{},
		&AuditLog{},
	)
}
