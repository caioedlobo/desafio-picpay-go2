package user

import (
	"context"
	"database/sql"
	"errors"
)

var ErrUserAlreadyExists = errors.New("user already exists")

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r Repository) Save(ctx context.Context, req *User) error {
	query := `INSERT INTO users (id, name, document_number, document_type, email, password_hash, created_at, balance)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
					RETURNING id`

	err := r.db.QueryRowContext(
		ctx,
		query,
		req.ID,
		req.Name,
		req.DocumentNumber,
		req.DocumentType,
		req.Email,
		req.Password.GetHash(),
		req.CreatedAt,
		req.Balance,
	).Scan(&req.ID)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrUserAlreadyExists
		default:
			return err
		}
	}
	return nil
}

func (r Repository) FindByEmail(ctx context.Context, email string) (*User, error) {
	query := `
        SELECT id, name, document_number, document_type, email, password_hash, created_at
        FROM users
        WHERE email = $1`
	var u User
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&u.ID,
		&u.Name,
		&u.DocumentNumber,
		&u.DocumentType,
		&u.Email,
		&u.Password,
		&u.CreatedAt,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, nil
		default:
			return nil, err
		}
	}
	return &u, nil
}
