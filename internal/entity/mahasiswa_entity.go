package entity

import "github.com/google/uuid"

type Mahasiswa struct {
	Timestamp
	Id       uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	GugusID  uuid.UUID `json:"gugus_id"`
	RegionID uuid.UUID `json:"region_id"`
	Name     string    `gorm:"type:varchar(255);not null" json:"name"`
	NRP      string    `gorm:"type:varchar(100);not null;unique" json:"NRP"`

	Gugus  *Gugus  `gorm:"foreignKey:GugusID" json:"gugus,omitempty"`
	Region *Region `gorm:"foreignKey:RegionID" json:"region,omitempty"`
}
