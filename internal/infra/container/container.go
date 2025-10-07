package container

import (
	"context"
	"database/sql"
	"desafio-picpay-go2/internal/domain/user"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	"os"
)

type Container struct {
	db             *sql.DB
	UserService    user.Service
	UserRepository user.UserRepository
}

func NewContainer(ctx context.Context) *Container {
	container := &Container{}
	container.initInfra(ctx)
	container.initRepositories()
	container.initServices()
}

func (c *Container) initInfra(ctx context.Context) {
	db, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	c.db = db
}
func (c *Container) initRepositories() {
	c.UserRepository = user.NewRepository(c.db)
}

func (c *Container) initServices() {
	c.UserService = *user.NewService(c.UserRepository)
}
