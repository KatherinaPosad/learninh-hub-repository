package reviews

import (
	"context"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetReviews(ctx context.Context) ([]Review, error) {
	return s.repository.GetReviews(ctx)
}

func (s *Service) CreateReview(ctx context.Context, review *Review) error {
	return s.repository.CreateReview(ctx, review)
}

func (s *Service) DeleteReview(ctx context.Context, review Review) (int64, error) {
	result, err := s.repository.DeleteReview(ctx, review)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil //для проверки, действительно ли отзыв был удалён
}
