package application

import (
	"github.com/joaofilippe/subclub/internal/adapter/repository"
	services "github.com/joaofilippe/subclub/internal/adapter/service"
	"github.com/joaofilippe/subclub/internal/domain/user"
	"github.com/joaofilippe/subclub/internal/infra/database"
	"github.com/joaofilippe/subclub/internal/infra/server"
)

// Application orquestra os componentes principais do sistema
type Application struct {
	dbConnection     *database.Connection
	server *server.Server
	userService user.Service
}

// New inicializa uma nova Application com suas dependências globais
func New(server *server.Server, db *database.Connection) *Application {
	return &Application{
		dbConnection:     db,
		server: server,
	}
}

func (a *Application) InitServices() error {
	userRepo := repository.NewUserPostgresRepository(a.dbConnection.GetDB())
	a.userService = services.NewUserService(userRepo)
	return nil
}

func (a *Application) Start(port string) error {
	return a.server.Start(port)
}


