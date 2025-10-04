package user

import (
	"context"
	"database/sql"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db: db}
}

func (r repository) Save(ctx context.Context, req *User) error {
	query := `INSERT INTO users (id, name, document_number, document_type, email, password, created_at, balance)
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

	return err
}
