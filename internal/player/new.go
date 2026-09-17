package player

import (
	"fmt"
	"net"
	"os/exec"
	"time"
)

type Player struct {
	cmd  *exec.Cmd
	conn net.Conn
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
	return &Player{cmd: cmd, conn: conn}, nil
}

func connectIPC(sockPath string) (net.Conn, error) {
	for i := 0; i < 50; i++ {
		conn, err := net.Dial("unix", sockPath)
		if err == nil {
			return conn, nil
		}
		time.Sleep(50 * time.Millisecond)
	}
	return nil, fmt.Errorf("mpv IPC socket %s never appeared after 2.5s", sockPath)
}

func (p *Player) Close() error {
	_ = p.Quit()
	_ = p.conn.Close()
	return p.cmd.Wait()
}
