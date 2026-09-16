package player

import "encoding/json"

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
	return err
}

func (p *Player) Play(url string) error {
	return p.sendCommand("loadfile", url, "replace")
}
