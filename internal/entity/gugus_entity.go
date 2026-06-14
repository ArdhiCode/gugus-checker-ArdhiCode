package entity

import "github.com/google/uuid"

type Gugus struct {
	Timestamp
	Id    uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Gugus string    `gorm:"type:varchar(255);not null" json:"gugus"`
}
