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
)

type Event struct {
	Type EventType
}

type jsonCheckEvent struct {
	Event string `json:"event"`
}

type jsonEndFile struct {
	Reason string `json:"reason"`
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
		log.Printf("%s %s", "READ: ", string(line))
		p.handleLine(line)
	}
}

func (p *Player) handleLine(line []byte) {
	var check jsonCheckEvent
	err := json.Unmarshal(line, &check)
	if err != nil {
		log.Printf("player.handleLine(): %s %s", err.Error(), string(line))
	}

	switch check.Event {
	case "end-file":
		p.endFile(line)
	}
}

func (p *Player) endFile(line []byte) {
	var end jsonEndFile
	err := json.Unmarshal(line, &end)
	if err != nil {
		log.Printf("player.endFile(): %s %s", err.Error(), string(line))
	}
	switch end.Reason {
	case "eof":
		p.onEvent(Event{Finished})
	case "stop":
		p.onEvent(Event{Stopped})
	}

}
