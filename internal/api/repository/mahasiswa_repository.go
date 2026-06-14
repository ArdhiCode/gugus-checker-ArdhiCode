package repository

import (
	"context"

	"github.com/ArdhiCode/gugus-checker-ArdhiCode/internal/entity"
	"gorm.io/gorm"
)

type MahasiswaRepository interface {
	GetByNRP(ctx context.Context, tx *gorm.DB, nrp string) (entity.Mahasiswa, error)
}

type mahasiswaRepository struct {
	db *gorm.DB
}

func NewMahasiswa(db *gorm.DB) MahasiswaRepository {
	return &mahasiswaRepository{db}
}

func (r *mahasiswaRepository) GetByNRP(ctx context.Context, tx *gorm.DB, nrp string) (entity.Mahasiswa, error) {
	db := r.db
	if tx != nil {
		db = tx
	}

	var result entity.Mahasiswa
	if err := db.WithContext(ctx).Where("nrp = ?", nrp).Preload("Gugus").Preload("Region").First(&result).Error; err != nil {
		return entity.Mahasiswa{}, nil
	}

	return result, nil
}
