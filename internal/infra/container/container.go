package container

import (
	"desafio-picpay-go2/internal/config"
	"desafio-picpay-go2/internal/domain/user"
	"desafio-picpay-go2/internal/infra/database/pg"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
)

type Container struct {
	config         *config.Config
	db             *pg.Database
	UserService    user.Service
	UserRepository user.UserRepository
	Validator      *validator.Validate
}

func NewContainer(cfg *config.Config) (*Container, error) {
	container := &Container{
		config:    cfg,
		Validator: validator.New(validator.WithRequiredStructEnabled()),
	}
	err := container.initInfra()
	if err != nil {
		return nil, err
	}
	container.initRepositories()
	container.initServices()
	return container, nil
}

func (c *Container) initInfra() error {
	db, err := pg.NewConnection(c.config.DriverName, c.config.PostgresDSN)
	if err != nil {
		return err
	}
	c.db = db
	return nil
}
func (c *Container) initRepositories() {
	c.UserRepository = user.NewRepository(c.db.DB())
}

func (c *Container) initServices() {
	c.UserService = *user.NewService(c.UserRepository)
}
