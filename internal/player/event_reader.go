package player

import (
	"bufio"
	"encoding/json"
	"log"
)

type EventType int

const (
	Finished EventType = iota
	Stopped
	TimePos
)

type Event struct {
	Type EventType
	Time float64
}

type jsonEvent struct {
	Event  string          `json:"event"`
	Reason string          `json:"reason"`
	Name   string          `json:"name"`
	Data   json.RawMessage `json:"data"`
}

func (p *Player) emit(ev Event) {
	if p.onEvent != nil {
		p.onEvent(ev)
	}
}

func (p *Player) readLoop() {
	scanner := bufio.NewScanner(p.conn)
	for scanner.Scan() {
		line := scanner.Bytes()
		//log.Printf("%s %s", "READ: ", string(line))
		p.handleLine(line)
	}
}

func (p *Player) handleLine(line []byte) {
	var check jsonEvent
	err := json.Unmarshal(line, &check)
	if err != nil {
		log.Printf("player.handleLine(): %s %s", err.Error(), string(line))
	}

	switch check.Event {
	case "end-file":
		p.endFile(check)
	case "property-change":
		p.propertyChange(check)
	}
}

func (p *Player) endFile(check jsonEvent) {
	switch check.Reason {
	case "eof":
		p.onEvent(Event{Type: Finished})
	case "stop":
		p.onEvent(Event{Type: Stopped})
	}
}

func (p *Player) propertyChange(check jsonEvent) {
	switch check.Name {
	case "time-pos":
		var tp float64
		err := json.Unmarshal(check.Data, &tp)
		if err != nil {
			return
		}
		p.onEvent(Event{
			Type: TimePos,
			Time: tp,
		})
	}
}
