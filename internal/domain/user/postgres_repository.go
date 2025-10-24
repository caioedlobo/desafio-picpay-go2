package user

import (
	"context"
	"database/sql"
	"desafio-picpay-go2/internal/infra/database/model"
	"desafio-picpay-go2/pkg/fault"
	"errors"
)

var ErrRecordNoFound = errors.New("record not found")

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
			return fault.New("user already exists2", fault.WithError(err))
		default:
			return err
		}
	}
	return nil
}

func (r Repository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	query := `
        SELECT id, name, document_number, document_type, email, password_hash, created_at
        FROM users
        WHERE email = $1`

	var u model.User
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&u.ID,
		&u.Name,
		&u.DocumentNumber,
		&u.DocumentType,
		&u.Email,
		&u.PasswordHash,
		&u.CreatedAt,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNoFound
		default:
			return nil, err
		}
	}
	return &u, nil
}

func (r Repository) Login(ctx context.Context, email, password string) (*model.User, error) {
	query := `
        SELECT id, name, document_number, document_type, email, password_hash, created_at
        FROM users
        WHERE email = $1 and password_hash = $2`
	var u model.User
	err := r.db.QueryRowContext(ctx, query, email, password).Scan(
		&u.ID,
		&u.Name,
		&u.DocumentNumber,
		&u.DocumentType,
		&u.Email,
		&u.PasswordHash,
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

func (r Repository) FindByID(ctx context.Context, userID string) (*model.User, error) {
	query := `select name, document_number, (balance).number, email from users where id = $1`
	var u model.User
	err := r.db.QueryRowContext(ctx, query, userID).Scan(
		&u.Name,
		&u.DocumentNumber,
		&u.BalanceNumber,
		&u.Email,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNoFound
		default:
			return nil, err
		}
	}
	return &u, nil
}
