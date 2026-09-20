package controller

import (
	"log"

	"github.com/joshw34/navitui/internal/types"
)

func (c *Controller) PlaySong(s types.Song) error {
	url, err := c.srv.GetStreamURL(s.ID)
	if err != nil {
		return err
	}
	err = c.play.PlaySong(url)
	if err != nil {
		return err
	}
	c.setNowPlaying(s)
	return nil
}

func (c *Controller) TogglePause() error {
	return c.play.TogglePause()
}

func (c *Controller) Stop() error {
	c.setNowPlaying(types.Song{})
	return c.play.Stop()
}

func (c *Controller) setNowPlaying(s types.Song) {
	c.mu.Lock()
	c.nowPlaying = s
	log.Printf("nowPlaying: %s", c.nowPlaying.Title)
	c.mu.Unlock()
	c.UIUpdate(Update{Type: NowPlaying, NowPlaying: s})
}
