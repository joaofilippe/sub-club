package api

import (
	userhandler "github.com/joaofilippe/subclub/internal/adapter/api/user"
	"github.com/joaofilippe/subclub/internal/application"
	"github.com/joaofilippe/subclub/internal/infra/server"
)

type API struct {
	server *server.Server
	application *application.Application
	userHandler *userhandler.UserHandler
}

func NewAPI(
	server *server.Server,
	application *application.Application,
	userHandler *userhandler.UserHandler,
) *API {
	return &API{
		server: server,
		application: application,
		userHandler: userHandler,
	}
}

// RegisterRoutes registra todas as rotas da API
func (a *API) RegisterRoutes() {
	a.registerUserRoutes()
}

// registerUserRoutes registra as rotas de usuário
func (a *API) registerUserRoutes() {
	server := a.server.GetEcho()
	userGroup := server.Group("/users")
	userGroup.POST("/", a.userHandler.Create)
}

// Start inicia todos os processos da aplicação (ex: o servidor HTTP)
func (a *API) Start(port string) error {
	return a.server.Start(":" + port)
}