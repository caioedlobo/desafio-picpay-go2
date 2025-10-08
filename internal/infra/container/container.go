package container

import (
	"database/sql"
	"desafio-picpay-go2/internal/config"
	"desafio-picpay-go2/internal/domain/user"
	_ "github.com/lib/pq"
)

type Container struct {
	config         *config.Config
	db             *sql.DB
	UserService    user.Service
	UserRepository user.UserRepository
}

func NewContainer(cfg *config.Config) *Container {
	container := &Container{
		config: cfg,
	}
	container.initInfra()
	container.initRepositories()
	container.initServices()
	return container
}

func (c *Container) initInfra() {
	db, err := sql.Open(c.config.DriverName, c.config.PostgresDSN)
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
