package controller

import (
	"slices"

	"github.com/joshw34/navitui/internal/types"
)

func (c *Controller) QueueAddToEnd(s types.Song) {
	c.mu.Lock()
	c.queue = slices.Insert(c.queue, len(c.queue), s)
	snapshot := slices.Clone(c.queue)
	c.mu.Unlock()
	c.updateUIQueue(snapshot)
}

func (c *Controller) QueueAddNext(s types.Song) {
	c.mu.Lock()
	c.queue = slices.Insert(c.queue, 0, s)
	snapshot := slices.Clone(c.queue)
	c.mu.Unlock()
	c.updateUIQueue(snapshot)
}

func (c *Controller) QueueRemove(i int) {
	c.mu.Lock()
	if i < 0 || i >= len(c.queue) {
		c.mu.Unlock()
		return
	}
	c.queue = slices.Delete(c.queue, i, i+1)
	snapshot := slices.Clone(c.queue)
	c.mu.Unlock()
	c.updateUIQueue(snapshot)
}

func (c *Controller) QueueClear() {
	c.mu.Lock()
	c.queue = slices.Delete(c.queue, 0, len(c.queue))
	snapshot := slices.Clone(c.queue)
	c.mu.Unlock()
	c.updateUIQueue(snapshot)
}

func (c *Controller) QueuePlayNext() {
	c.mu.Lock()
	if len(c.queue) == 0 {
		c.mu.Unlock()
		return
	}
	next := c.queue[0]
	c.queue = slices.Delete(c.queue, 0, 1)
	snapshot := slices.Clone(c.queue)
	c.mu.Unlock()
	_ = c.PlaySong(next)
	c.updateUIQueue(snapshot)
}

func (c *Controller) updateUIQueue(q []types.Song) {
	if c.UIUpdate != nil {
		c.UIUpdate(Update{Type: QueueUpdate, Queue: q})
	}
}
