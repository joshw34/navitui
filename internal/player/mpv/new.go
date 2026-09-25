package mpv

import (
	"net"
	"os/exec"
	"sync"

	"github.com/joshw34/navitui/internal/player"
	"github.com/joshw34/navitui/internal/types"
)

var _ types.Player = (*Mpv)(nil)

type Mpv struct {
	cmd  *exec.Cmd
	conn net.Conn
	pending
	onEvent func(player.Event) // Set during controller.New()
}

type pending struct {
	pendingPlays map[uint64]uint64
	waitingLoad  uint64
	mu           sync.Mutex
}

func New() (player *Mpv, err error) {
	cmd := exec.Command("mpv", "--idle", "--input-ipc-server="+ipcSocketPath, "--no-video")
	if err := cmd.Start(); err != nil {
		return nil, err
	}

	conn, err := connectIPC() // platform-specific, no arg
	if err != nil {
		return nil, err
	}
	p := &Mpv{cmd: cmd, conn: conn}
	p.startReader()
	p.pendingPlays = map[uint64]uint64{}
	err = p.observeTimePos()
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Mpv) SetEventHandler(fn func(player.Event)) {
	p.onEvent = fn
}

func (p *Mpv) startReader() {
	go p.readLoop()
}

func (p *Mpv) Close() {
	_ = p.Quit()
	_ = p.conn.Close()
	_ = p.cmd.Wait()
}
