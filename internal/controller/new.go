package controller

import (
	"github.com/joshw34/navitui/internal/cache"
	"github.com/joshw34/navitui/internal/client"
)

type Controller struct {
	srv *client.Client
	db  *cache.Cache
}

func New(srv *client.Client, db *cache.Cache) *Controller {
	return &Controller{
		srv: srv,
		db:  db,
	}
}
