package controller

import (
	"github.com/joshw34/navitui/internal/types"
)

func (c *Controller) QueueAddToEnd(s types.Song) {
	c.queueInsert(len(c.queue), s)
}

func (c *Controller) QueueAddNext(s types.Song) {
	c.queueInsert(0, s)
}

func (c *Controller) QueueRemove(i int) {
	c.queueDelete(i, i+1)
}

func (c *Controller) QueueClear() {
	c.queueDelete(0, len(c.queue))
}

func (c *Controller) QueueAdvance() error {
	if c.queueIsEmpty() {
		return nil
	}
	next := c.queuePopNext()
	return c.PlaySong(next)
}
