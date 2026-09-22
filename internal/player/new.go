package player

import (
	"net"
	"os/exec"
	"sync"
)

type Player struct {
	cmd  *exec.Cmd
	conn net.Conn
	pending
	onEvent func(Event) // Set during controller.New()
}

type pending struct {
	pendingPlays map[uint64]uint64
	waitingLoad  uint64
	mu           sync.Mutex
}

func New() (player *Player, err error) {
	cmd := exec.Command("mpv", "--idle", "--input-ipc-server="+ipcSocketPath, "--no-video")
	if err := cmd.Start(); err != nil {
		return nil, err
	}

	conn, err := connectIPC() // platform-specific, no arg
	if err != nil {
		return nil, err
	}
	p := &Player{cmd: cmd, conn: conn}
	p.startReader()
	p.pendingPlays = map[uint64]uint64{}
	err = p.observeTimePos()
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Player) SetEventHandler(fn func(Event)) {
	p.onEvent = fn
}

func (p *Player) startReader() {
	go p.readLoop()
}

func (p *Player) Close() error {
	_ = p.Quit()
	_ = p.conn.Close()
	return p.cmd.Wait()
}
