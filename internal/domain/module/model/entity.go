package model

import "time"

type Module struct {
	ID        string
	Name      string
	Active    bool
	CreatedAt time.Time
}
