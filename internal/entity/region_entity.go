package entity

import "github.com/google/uuid"

type Region struct {
	Timestamp
	Id     uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Region string    `gorm:"type:varchar(255);not null" json:"region"`
}
