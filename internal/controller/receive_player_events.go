package controller

import (
	"github.com/joshw34/navitui/internal/player"
	"github.com/joshw34/navitui/internal/types"
)

type UpdateType int

const (
	NowPlaying UpdateType = iota
	QueueUpdate
	TimePosUpdate
)

type Update struct {
	Type       UpdateType
	NowPlaying types.Song
	Queue      []types.Song
	TimePos    float64
}

func (c *Controller) PlayerEventHandler(e player.Event) {
	switch e.Type {
	case player.Finished:
		c.uiTimePos(0)
		c.QueueAdvance()
	case player.Stopped:
		c.uiTimePos(0)
		return
	case player.TimePos:
		c.uiTimePos(e.Time)
	}
}
