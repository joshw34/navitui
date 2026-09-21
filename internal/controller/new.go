// Package controller: this packages coordinates the other packages
package controller

import (
	"sync"
	"sync/atomic"

	"github.com/joshw34/navitui/internal/cache"
	"github.com/joshw34/navitui/internal/client"
	"github.com/joshw34/navitui/internal/player"
	"github.com/joshw34/navitui/internal/types"
)

type Controller struct {
	srv  *client.Client
	db   *cache.Cache
	play *player.Player
	PlayerState
	UIUpdate func(Update)
}

type PlayerState struct {
	queue           []types.Song
	queueMutex      sync.Mutex
	pendingPlays    map[uint64]types.Song
	pendingMutex    sync.Mutex
	nowPlaying      types.Song
	nowPlayingMutex sync.Mutex
	timePos         float64
	timePosMutex    sync.Mutex
	reqID           atomic.Uint64
	reqIDMutex      sync.Mutex
}

func New(srv *client.Client, db *cache.Cache, play *player.Player) *Controller {
	c := &Controller{
		srv:          srv,
		db:           db,
		play:         play,
		queue:        []types.Song{},
		pendingPlays: map[uint64]types.Song{},
		nowPlaying:   types.Song{},
		timePos:      -1,
	}
	c.play.SetEventHandler(c.PlayerEventHandler)
	return c
}

func (c *Controller) SetUpdateHandler(f func(u Update)) {
	c.UIUpdate = f
}
