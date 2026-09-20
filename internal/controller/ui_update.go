package controller

import (
	"math"

	"github.com/joshw34/navitui/internal/types"
)

func (c *Controller) uiQueue(q []types.Song) {
	if c.UIUpdate != nil {
		c.UIUpdate(Update{Type: QueueUpdate, Queue: q})
	}
}

func (c *Controller) uiTimePos(tp float64) {
	c.mu.Lock()
	lastSent := c.timePos
	send := math.Abs(tp-lastSent) >= 0.5
	if send {
		c.timePos = tp
	}
	c.mu.Unlock()
	if send && c.UIUpdate != nil {
		c.UIUpdate(Update{Type: TimePosUpdate, TimePos: tp})
	}
}
