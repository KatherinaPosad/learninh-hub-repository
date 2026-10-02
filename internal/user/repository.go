package user

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

/* описывает sql запрос и отправляет это запрос к БД в БД
func CreateTable(ctx context.Context, conn *pgx.Conn) error{
sqlQuery := `
CREATE TABLE IF NOT EXISTS users
и тут описание столбцов
);
`
_, err:=conn.Exec(ctx, sqlQuery)
if err!=nil{
return err
}
return nil
}
*/

func (r *Repository) CreateUser(ctx context.Context, user *User) error {
	err := r.db.QueryRow(ctx,
		`
		INSERT INTO users (login, name, email, password_hash, role)
		VALUES (NULLIF($1, ''), NULLIF($2, ''), NULLIF($3, ''), $4, $5)
		RETURNING id
		`,
		user.Login,
		user.Name,
		user.Email,
		user.PasswordHash,
		user.Role,
	).Scan(&user.ID)

	return err
}

func (r *Repository) GetUserByLogin(ctx context.Context, login string) (User, error) {
	var user User

	err := r.db.QueryRow(ctx, //если login = NULL вернуть вместо него ""
		`
		SELECT 
			id, 
			COALESCE (login, ''),
			COALESCE (name, ''),
			COALESCE (email, ''),
			password_hash,
			role
		FROM users
		WHERE login = $1
		`,
		login,
	).Scan(
		&user.ID,
		&user.Login,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
	)
	return user, err
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (User, error) {
	var user User

	err := r.db.QueryRow(ctx,
		`
			SELECT 
			id, 
			COALESCE (login, ''),
			COALESCE (name, ''),
			COALESCE (email, ''),
			password_hash,
			role
		FROM users
		WHERE email = $1
		`,
		email,
	).Scan(
		&user.ID,
		&user.Login,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
	)
	return user, err
}
