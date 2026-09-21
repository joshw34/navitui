package controller

import (
	"slices"

	"github.com/joshw34/navitui/internal/types"
)

type SeekDirection int

const (
	Forward SeekDirection = iota
	Backward
)

func (c *Controller) PlaySong(s types.Song) error {
	url, err := c.srv.GetStreamURL(s.ID)
	if err != nil {
		return err
	}
	reqID := c.getReqID()
	err = c.play.PlaySong(reqID, url)
	if err != nil {
		return err
	}
	c.pendingAdd(reqID, s)
	return nil
}

func (c *Controller) TogglePlayPause() error {
	if c.trackIsPlaying() {
		return c.play.TogglePause()
	}
	if c.queueIsEmpty() {
		return nil
	}
	next := c.queuePopNext()
	err := c.PlaySong(next)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) Stop() error {
	err := c.play.Stop()
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) Seek(d SeekDirection) error {
	if !c.trackIsPlaying() {
		return nil
	}
	pos := c.getTimePos()
	var newPos float64
	switch d {
	case Forward:
		newPos = pos + 5
	case Backward:
		newPos = pos - 5
	}
	return c.play.Seek(newPos)
}

func (c *Controller) Skip(d SeekDirection) error {
	if !c.trackIsPlaying() {
		return nil
	}
	switch d {
	case Forward:
		if c.queueIsEmpty() {
			return c.play.Stop()
		}
		next := c.queuePopNext()
		return c.PlaySong(next)
	case Backward:
		return c.play.Seek(0)
	}
	return nil
}

// MUTEX-PROTECTED HELPERS

func (c *Controller) setNowPlaying(s types.Song) {
	c.nowPlayingMutex.Lock()
	defer c.nowPlayingMutex.Unlock()
	c.nowPlaying = s
	c.uiNowPLaying(s)
}

func (c *Controller) getReqID() uint64 {
	c.reqIDMutex.Lock()
	defer c.reqIDMutex.Unlock()
	id := c.reqID.Add(1)
	if id == 0 {
		id = c.reqID.Add(1)
	}
	return id
}

func (c *Controller) getTimePos() float64 {
	c.timePosMutex.Lock()
	defer c.timePosMutex.Unlock()
	return c.timePos
}

func (c *Controller) pendingAdd(reqID uint64, s types.Song) {
	c.pendingMutex.Lock()
	defer c.pendingMutex.Unlock()
	c.pendingPlays[reqID] = s
}

func (c *Controller) pendingRemove(reqID uint64) types.Song {
	c.pendingMutex.Lock()
	defer c.pendingMutex.Unlock()
	s := c.pendingPlays[reqID]
	delete(c.pendingPlays, reqID)
	return s
}

func (c *Controller) queuePopNext() types.Song {
	c.queueMutex.Lock()
	defer c.queueMutex.Unlock()
	next := c.queue[0]
	c.queue = slices.Delete(c.queue, 0, 1)
	c.uiQueue(slices.Clone(c.queue))
	return next
}

func (c *Controller) queueInsert(i int, s types.Song) {
	c.queueMutex.Lock()
	defer c.queueMutex.Unlock()
	c.queue = slices.Insert(c.queue, i, s)
	c.uiQueue(slices.Clone(c.queue))
}

func (c *Controller) queueDelete(start, end int) {
	c.queueMutex.Lock()
	if start < 0 || start >= len(c.queue) {
		return
	}
	defer c.queueMutex.Unlock()
	c.queue = slices.Delete(c.queue, start, end)
	c.uiQueue(slices.Clone(c.queue))
}

func (c *Controller) queueIsEmpty() bool {
	c.queueMutex.Lock()
	defer c.queueMutex.Unlock()
	return len(c.queue) == 0
}

func (c *Controller) trackIsPlaying() bool {
	c.nowPlayingMutex.Lock()
	defer c.nowPlayingMutex.Unlock()
	return len(c.nowPlaying.ID) > 0
}
