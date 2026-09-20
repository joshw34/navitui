package controller

import (
	"github.com/joshw34/navitui/internal/player"
	"github.com/joshw34/navitui/internal/types"
)

type UpdateType int

const (
	NowPlaying UpdateType = iota
	QueueUpdate
)

type Update struct {
	Type       UpdateType
	NowPlaying types.Song
	Queue      []types.Song
}

func (c *Controller) PlayerEventHandler(e player.Event) {
	switch e.Type {
	case player.Finished:
		c.QueuePlayNext()
	case player.Stopped:
		return
	}
}
