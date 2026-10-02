package category

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

func (s *Service) GetCategories(ctx context.Context) ([]Category, error) {
	return s.repository.GetCategories(ctx)
}

func (s *Service) CreateCategory(ctx context.Context, category *Category) error {
	if category.Name == "" {
		return errors.New("categiry name must not be empty")
	}

	return s.repository.CreateCategory(ctx, category)
}
