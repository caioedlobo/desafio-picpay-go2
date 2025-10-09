package pg

import (
	"context"
	"database/sql"
	"time"
)

type Database struct {
	db *sql.DB
}

func NewConnection(driverName, dsn string) (*Database, error) {
	db, err := sql.Open(driverName, dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

func (d *Database) Close() error {
	return d.db.Close()
}

func (d *Database) DB() *sql.DB {
	return d.db
}
