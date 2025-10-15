package container

import (
	"desafio-picpay-go2/internal/config"
	"desafio-picpay-go2/internal/domain/user"
	"desafio-picpay-go2/internal/infra/database/pg"
	"github.com/charmbracelet/log"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
)

type Container struct {
	logger         *log.Logger
	Config         *config.Config
	db             *pg.Database
	UserService    user.UserService
	UserRepository user.UserRepository
	Validator      *validator.Validate
}

func NewContainer(cfg *config.Config, log *log.Logger) (*Container, error) {
	container := &Container{
		Config:    cfg,
		logger:    log,
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
	db, err := pg.NewConnection(c.Config.DriverName, c.Config.PostgresDSN)
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
	c.UserService = *user.NewService(c.UserRepository, c.logger, c.Config.JWTSecretKey, c.Config.JWTAccessTokenDuration)
}
