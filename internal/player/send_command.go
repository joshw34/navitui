package player

import (
	"encoding/json"
	"log"
)

type ipcCommand struct {
	Command   []any `json:"command"`
	RequestID []int `json:"request_id,omitempty"`
}

func (p *Player) sendCommand(args ...any) error {
	cmd := ipcCommand{Command: args}
	data, err := json.Marshal(cmd)
	if err != nil {
		return err
	}
	data = append(data, '\n')
	_, err = p.conn.Write(data)
	log.Printf("%s %s", "SEND: ", string(data))
	return err
}

func (p *Player) PlaySong(url string) error {
	return p.sendCommand("loadfile", url, "replace")
}

func (p *Player) Stop() error {
	return p.sendCommand("stop")
}

func (p *Player) TogglePause() error {
	return p.sendCommand("cycle", "pause")
}

func (p *Player) Quit() error {
	return p.sendCommand("quit")
}
