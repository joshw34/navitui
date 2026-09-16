package controller

func (c *Controller) Play(songID string) error {
	url, err := c.srv.GetStreamURL(songID)
	if err != nil {
		return err
	}
	return c.Play(url)
}
