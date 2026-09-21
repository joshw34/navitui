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
	case player.PlaybackEOF:
		c.uiTimePos(0)
		c.setNowPlaying(types.Song{})
		_ = c.QueueAdvance()
	case player.PlaybackStopped:
		c.uiTimePos(0)
		c.setNowPlaying(types.Song{})
	case player.PlaybackPos:
		c.uiTimePos(e.Time)
	case player.PlaybackStarted:
		c.setNowPlaying(c.pendingRemove(e.RequestId))
	case player.PlaybackError:
		_ = c.pendingRemove(e.RequestId)
	}
}
