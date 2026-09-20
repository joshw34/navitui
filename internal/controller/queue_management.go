package controller

import (
	"slices"

	"github.com/joshw34/navitui/internal/types"
)

func (c *Controller) QueueAddToEnd(s types.Song) {
	c.mu.Lock()
	snapshot := c.insertAndReturnQueue(len(c.queue), s)
	c.mu.Unlock()
	c.uiQueue(snapshot)
}

func (c *Controller) QueueAddNext(s types.Song) {
	c.mu.Lock()
	snapshot := c.insertAndReturnQueue(0, s)
	c.mu.Unlock()
	c.uiQueue(snapshot)
}

func (c *Controller) QueueRemove(i int) {
	c.mu.Lock()
	if i < 0 || i >= len(c.queue) {
		c.mu.Unlock()
		return
	}
	snapshot := c.deleteAndReturnQueue(i, i+1)
	c.mu.Unlock()
	c.uiQueue(snapshot)
}

func (c *Controller) QueueClear() {
	c.mu.Lock()
	snapshot := c.deleteAndReturnQueue(0, len(c.queue))
	c.mu.Unlock()
	c.uiQueue(snapshot)
}

func (c *Controller) QueueAdvance() {
	c.mu.Lock()
	if c.queueEmpty() {
		c.mu.Unlock()
		return
	}
	next, snapshot := c.popNextAndReturnQueue()
	c.mu.Unlock()
	c.uiQueue(snapshot)
	_ = c.PlaySong(next)
}

// These functions must only be called when the mutex is locked

func (c *Controller) popNextAndReturnQueue() (types.Song, []types.Song) {
	next := c.queue[0]
	return next, c.deleteAndReturnQueue(0, 1)
}

func (c *Controller) insertAndReturnQueue(i int, s types.Song) []types.Song {
	c.queue = slices.Insert(c.queue, i, s)
	return slices.Clone(c.queue)
}

func (c *Controller) deleteAndReturnQueue(start, end int) []types.Song {
	c.queue = slices.Delete(c.queue, start, end)
	return slices.Clone(c.queue)
}

func (c *Controller) queueEmpty() bool {
	return len(c.queue) == 0
}

func (c *Controller) trackPlaying() bool {
	return len(c.nowPlaying.ID) > 0
}
