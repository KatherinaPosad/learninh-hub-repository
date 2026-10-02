package reviews

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Repository struct {
	db *pgx.Conn
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetReviews(ctx context.Context) ([]Review, error) {
	rows, err := r.db.Query(ctx,
		`
SELECT id, user_id, material_id, rating, comment, created_at 
FROM reviews
ORDER BY created_at

`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []Review
	for rows.Next() {
		var review Review
		err := rows.Scan(
			&review.ID,
			&review.UserID,
			&review.MaterialID,
			&review.Rating,
			&review.Comment,
			&review.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}

	return reviews, err

}

func (r *Repository) CreateReview(ctx context.Context, review *Review) error {
	err := r.db.QueryRow(ctx,
		`
	INSERT INTO reviews (user_id, material_id, rating, comment, created_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id
	`,
		review.UserID,
		review.MaterialID,
		review.Rating,
		review.Comment,
		review.CreatedAt,
	).Scan(&review.ID)
	return err
}

func (r *Repository) DeleteReview(ctx context.Context, review Review) (pgconn.CommandTag, error) {
	result, err := r.db.Exec(ctx,
		`
	DELETE FROM reviews
	WHERE id = $1 AND user_id = $2
	`,
		review.ID,
		review.UserID,
	)

	return result, err
}
