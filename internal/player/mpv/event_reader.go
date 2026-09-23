package mpv

import (
	"bufio"
	"encoding/json"
	"log"

	"github.com/joshw34/navitui/internal/player"
)

type jsonEvent struct {
	RequestID       uint64          `json:"request_id"`
	Event           string          `json:"event"`
	Reason          string          `json:"reason"`
	Name            string          `json:"name"`
	Error           string          `json:"error"`
	PlaylistEntryID uint64          `json:"playlist_entry_id"`
	Data            json.RawMessage `json:"data"`
}

type jsonPlaylistEntryID struct {
	PlaylistEntryID uint64 `json:"playlist_entry_id"`
}

func (p *Mpv) emit(ev player.Event) {
	if p.onEvent != nil {
		p.onEvent(ev)
	}
}

func (p *Mpv) readLoop() {
	scanner := bufio.NewScanner(p.conn)
	for scanner.Scan() {
		line := scanner.Bytes()
		var logcheck jsonEvent
		err := json.Unmarshal(line, &logcheck)
		if err != nil {
			log.Println("error unmarshalling event:", err)
		}
		if logcheck.Name != "time-pos" {
			log.Printf("%s %s", "READ: ", string(line))
		}
		p.handleLine(line)
	}
}

func (p *Mpv) handleLine(line []byte) {
	var check jsonEvent
	err := json.Unmarshal(line, &check)
	if err != nil {
		log.Printf("player.handleLine(): %s %s", err.Error(), string(line))
	}
	if check.RequestID >= player.MinReqId ||
		check.Event == "start-file" ||
		check.Event == "file-loaded" ||
		check.Event == "end-file" {
		if p.playStatus(check) {
			return
		}
	}
	if check.Event == "property-change" {
		p.propertyChange(check)
	}
}

func (p *Mpv) propertyChange(check jsonEvent) {
	switch check.Name {
	case "time-pos":
		var tp float64
		err := json.Unmarshal(check.Data, &tp)
		if err != nil {
			return
		}
		p.emit(player.Event{
			Type: player.PlaybackPos,
			Time: tp,
		})
	}
}

func (p *Mpv) playStatus(check jsonEvent) bool {
	var pID jsonPlaylistEntryID
	pIDErr := json.Unmarshal(check.Data, &pID)
	switch {
	case pIDErr == nil && check.Error == "success": // loadfile accepted (playback may still fail)
		p.pendingPlays[pID.PlaylistEntryID] = check.RequestID
		return true

	case check.Event == "start-file": // file will be next to load (err dealt with in endFile)
		p.waitingLoad = check.PlaylistEntryID
		return true

	case check.Event == "file-loaded": // file from last start-file has begun playing. clear from pending and waiting load
		reqID := p.pendingPlays[p.waitingLoad]
		if reqID < player.MinReqId { // missing requestId
			return true
		}
		delete(p.pendingPlays, p.waitingLoad)
		p.waitingLoad = 0
		p.emit(player.Event{Type: player.PlaybackStarted, RequestId: reqID})
		return true
	}

	if check.Event == "end-file" {
		switch check.Reason {
		case "error":
			reqID := p.pendingPlays[check.PlaylistEntryID]
			if reqID < player.MinReqId {
				return true
			}
			delete(p.pendingPlays, check.PlaylistEntryID)
			if p.waitingLoad == check.PlaylistEntryID {
				p.waitingLoad = 0
			}
			p.emit(player.Event{Type: player.PlaybackError, RequestId: reqID})
			return true
		case "eof":
			p.emit(player.Event{Type: player.PlaybackEOF})
			return true
		case "stop":
			p.emit(player.Event{Type: player.PlaybackStopped})
			return true
		}
	}
	return false
}
