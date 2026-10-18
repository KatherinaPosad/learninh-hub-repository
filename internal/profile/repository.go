package profile

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	db *pgx.Conn
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) UpdateProfile(ctx context.Context, userID int, request UpdateProfileRequest) (Profile, error) {
	var profile Profile

	err := r.db.QueryRow(ctx,
		`
		UPDATE users
		SET
			login = NULLIF($1, ''),
			name = NULLIF($2, ''),
			email = NULLIF($3, '')
		WHERE id = $4
		RETURNING
			id,
			COALESCE(login,''),
			COALESCE(name, ''),
			COALESCE(email, '')
		`,
		request.Login,
		request.Name,
		request.Email,
		userID,
	).Scan(
		&profile.ID,
		&profile.Login,
		&profile.Name,
		&profile.Email,
	)

	return profile, err
}
