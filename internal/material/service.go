package material

import (
	"context"
	"errors"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetMaterials(ctx context.Context) ([]Material, error) {
	return s.repository.GetMaterials(ctx)
}

func (s *Service) CreateMaterial(ctx context.Context, material *Material) error {
	if material.Title == "" {
		return errors.New("The title must be not empty")
	}

	if material.URL == "" {
		return errors.New("The URL must be not empty")
	}

	if material.Type == "" {
		return errors.New("The type must be not empty")
	}

	return s.repository.CreateMaterial(ctx, *material)
}
