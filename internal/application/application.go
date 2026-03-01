package application

import (
	"github.com/jmoiron/sqlx"
	"github.com/joaofilippe/subclub/internal/infra/server"
)

// Application orquestra os componentes principais do sistema (Servidor HTTP, Banco e futuros conectores)
type Application struct {
	server *server.Server
	db     *sqlx.DB
}

// New inicializa uma nova Application com suas dependências globais
func New(server *server.Server, db *sqlx.DB) *Application {
	return &Application{
		server: server,
		db:     db,
	}
}

// Start inicia todos os processos da aplicação (ex: o servidor HTTP)
func (a *Application) Start(port string) error {
	return a.server.Start(":" + port)
}
