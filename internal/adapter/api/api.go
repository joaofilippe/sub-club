package api

import (
	"github.com/joaofilippe/subclub/internal/application"
)

type API struct {
	application *application.Application
}

func NewAPI(application *application.Application) *API {
	return &API{
		application: application,
	}
}