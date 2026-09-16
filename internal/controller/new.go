package controller

import (
	"github.com/joshw34/navitui/internal/cache"
	"github.com/joshw34/navitui/internal/client"
	"github.com/joshw34/navitui/internal/player"
)

type Controller struct {
	srv  *client.Client
	db   *cache.Cache
	play *player.Player
}

func New(srv *client.Client, db *cache.Cache, play *player.Player) *Controller {
	return &Controller{
		srv:  srv,
		db:   db,
		play: play,
	}
}
