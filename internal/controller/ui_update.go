package controller

import (
	"math"

	"github.com/joshw34/navitui/internal/types"
)

// uiQueue - sends a clone, not the original
func (c *Controller) uiQueue(q []types.Song) {
	if c.UIUpdate != nil {
		c.UIUpdate(Update{Type: QueueUpdate, Queue: q})
	}
}

func (c *Controller) uiTimePos(tp float64) {
	c.timePosMutex.Lock()
	lastSent := c.timePos
	send := math.Abs(tp-lastSent) >= 0.5 || lastSent == 0
	if send {
		c.timePos = tp
	}
	c.timePosMutex.Unlock()
	if send && c.UIUpdate != nil {
		c.UIUpdate(Update{Type: TimePosUpdate, TimePos: tp})
	}
}

func (c *Controller) uiNowPlaying(s types.Song) {
	if c.UIUpdate != nil {
		c.UIUpdate(Update{Type: NowPlaying, NowPlaying: s})
	}
}
