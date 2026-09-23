// Package controller: this packages coordinates the other packages
package controller

import (
	"sync"
	"sync/atomic"

	"github.com/joshw34/navitui/internal/types"
)

type Controller struct {
	srv  types.Server
	db   types.Cache
	play types.Player
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

func New(srv types.Server, db types.Cache, play types.Player) *Controller {
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
