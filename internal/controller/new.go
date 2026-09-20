// Package controller: this packages coordinates the other packages
package controller

import (
	"sync"

	"github.com/joshw34/navitui/internal/cache"
	"github.com/joshw34/navitui/internal/client"
	"github.com/joshw34/navitui/internal/player"
	"github.com/joshw34/navitui/internal/types"
)

type Controller struct {
	srv  *client.Client
	db   *cache.Cache
	play *player.Player
	Queue
	UIUpdate func(Update)
}

type Queue struct {
	queue      []types.Song
	nowPlaying types.Song
	mu         sync.Mutex
}

func New(srv *client.Client, db *cache.Cache, play *player.Player) *Controller {
	c := &Controller{
		srv:        srv,
		db:         db,
		play:       play,
		queue:      []types.Song{},
		nowPlaying: types.Song{},
	}
	c.play.SetEventHandler(c.PlayerEventHandler)
	return c
}

func (c *Controller) SetUpdateHandler(f func(u Update)) {
	c.UIUpdate = f
}
