package controller

import (
	"github.com/joshw34/navitui/internal/types"
)

func (c *Controller) QueueAddToEnd(s types.Song) {
	c.queueAppend(s)
}

func (c *Controller) QueueAddNext(s types.Song) {
	c.queueInsert(0, s)
}

func (c *Controller) QueueRemove(i int) {
	c.queueDelete(i, i+1)
}

func (c *Controller) QueueClear() {
	c.queueDeleteAll()
}

func (c *Controller) QueueAdvance() error {
	next, ok := c.queuePopNext()
	if !ok {
		return nil // no error, queue is empty
	}
	return c.PlaySong(next)
}
