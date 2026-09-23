package mpv

import (
	"encoding/json"
	"log"
)

type ipcCommand struct {
	Command   []any  `json:"command"`
	RequestID uint64 `json:"request_id,omitempty"`
}

func (p *Mpv) sendCommand(reqId uint64, args ...any) error {
	cmd := ipcCommand{Command: args, RequestID: reqId}
	data, err := json.Marshal(cmd)
	if err != nil {
		return err
	}
	data = append(data, '\n')
	_, err = p.conn.Write(data)
	log.Printf("%s %s", "SEND: ", string(data))
	return err
}

func (p *Mpv) PlaySong(reqId uint64, url string) error {
	err := p.sendCommand(reqId, "loadfile", url, "replace")
	if err != nil {
		return err
	}
	return nil
}

func (p *Mpv) Stop() error {
	return p.sendCommand(0, "stop")
}

func (p *Mpv) TogglePause() error {
	return p.sendCommand(0, "cycle", "pause")
}

func (p *Mpv) Quit() error {
	return p.sendCommand(0, "quit")
}

func (p *Mpv) observeTimePos() error {
	return p.sendCommand(0, "observe_property", 0, "time-pos")
}

func (p *Mpv) Seek(t float64) error {
	return p.sendCommand(0, "set_property", "time-pos", t)
}
