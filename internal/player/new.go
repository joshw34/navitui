package player

import (
	"fmt"
	"net"
	"os/exec"
	"time"
)

type Player struct {
	cmd     *exec.Cmd
	conn    net.Conn
	onEvent func(Event) // Set during controller.New()
}

func New() (player *Player, err error) {
	cmd := exec.Command("mpv", "--idle", "--input-ipc-server=/tmp/navitui-mpv.sock", "--no-video")
	if err := cmd.Start(); err != nil {
		return nil, err
	}

	conn, err := connectIPC("/tmp/navitui-mpv.sock")
	if err != nil {
		return nil, err
	}
	p := &Player{cmd: cmd, conn: conn}
	p.startReader()
	return p, nil
}

func connectIPC(sockPath string) (net.Conn, error) {
	for range 50 {
		conn, err := net.Dial("unix", sockPath)
		if err == nil {
			return conn, nil
		}
		time.Sleep(50 * time.Millisecond)
	}
	return nil, fmt.Errorf("mpv IPC socket %s never appeared after 2.5s", sockPath)
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
