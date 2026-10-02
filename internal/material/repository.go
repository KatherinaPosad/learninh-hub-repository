package material

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

func (r *Repository) GetMaterials(ctx context.Context) ([]Material, error) {
	rows, err := r.db.Query(ctx,
		`
	SELECT id, title, description, url, category_id, type, created_at
	FROM materials
	ORDER BY id ASC
	`,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var materials []Material

	for rows.Next() {
		var material Material
		err := rows.Scan(
			&material.ID,
			&material.Title,
			&material.Description,
			&material.URL,
			&material.CategoryID,
			&material.Type,
			&material.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		materials = append(materials, material)
	}
	return materials, nil
}

func (r *Repository) CreateMaterial(ctx context.Context, material Material) error {
	err := r.db.QueryRow(ctx,
		`
	INSERT INTO materials (title, description, url, category_id, type, created_at)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id
	`,
		material.Title,
		material.Description,
		material.URL,
		material.CategoryID,
		material.Type,
		material.CreatedAt,
	).Scan(&material.ID)

	return err
}
