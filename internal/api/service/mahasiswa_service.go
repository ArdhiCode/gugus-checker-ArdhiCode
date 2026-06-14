package service

import (
	"context"
	"errors"

	"github.com/ArdhiCode/gugus-checker-ArdhiCode/internal/api/repository"
	"github.com/ArdhiCode/gugus-checker-ArdhiCode/internal/dto"
	"gorm.io/gorm"
)

type MahasiswaService interface {
	GetByNRP(ctx context.Context, req dto.GetByNRPRequest) (dto.GetByNRPResponse, error)
}

type mahasiswaService struct {
	mahasiswaRepo repository.MahasiswaRepository
}

func NewMahasiswa(mahasiswaRepo repository.MahasiswaRepository) MahasiswaService {
	return &mahasiswaService{mahasiswaRepo}
}
func (s *mahasiswaService) GetByNRP(ctx context.Context, req dto.GetByNRPRequest) (dto.GetByNRPResponse, error) {
	mahasiswa, err := s.mahasiswaRepo.GetByNRP(ctx, nil, req.NRP)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.GetByNRPResponse{}, errors.New("mahasiswa not found")
		}
		return dto.GetByNRPResponse{}, err
	}

	resp := dto.GetByNRPResponse{
		Name:   mahasiswa.Name,
		NRP:    mahasiswa.NRP,
		Gugus:  mahasiswa.Gugus.Gugus,
		Region: mahasiswa.Region.Region,
	}

	return resp, nil
}
