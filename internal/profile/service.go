package profile

import "context"

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) UpdateProfile(ctx context.Context, userID int, request UpdateProfileRequest) (Profile, error) {
	profile, err := s.repository.UpdateProfile(ctx, userID, request)
	if err != nil {
		return Profile{}, err
	}

	return profile, nil
}
