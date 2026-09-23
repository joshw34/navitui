package player

type EventType int

const (
	PlaybackEOF EventType = iota
	PlaybackStopped
	PlaybackPos
	PlaybackStarted
	PlaybackError
)

type Event struct {
	Type      EventType
	Time      float64
	RequestId uint64
}

const MinReqId = 1
