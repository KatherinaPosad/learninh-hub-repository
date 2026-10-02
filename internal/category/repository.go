package category

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

func (r *Repository) GetCategories(ctx context.Context) ([]Category, error) {
	rows, err := r.db.Query(ctx,
		`
	SELECT id, name
	FROM categories
	ORDER BY id
	`,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var categories []Category

	for rows.Next() {
		var category Category

		err := rows.Scan(
			&category.ID,
			&category.Name,
		)

		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}
	return categories, nil
}

func (r *Repository) CreateCategory(ctx context.Context, category *Category) error {
	err := r.db.QueryRow(ctx,
		`
		INSERT INTO categories (name) VALUES ($1) RETURNING id
		`,
		category.Name,
	).Scan(&category.ID)

	return err
}
