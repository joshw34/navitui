//go:build unix

package player

import (
	"fmt"
	"net"
	"time"
)

const ipcSocketPath = "/tmp/navitui-mpv.sock"

func connectIPC() (net.Conn, error) {
	for range 50 {
		conn, err := net.Dial("unix", ipcSocketPath)
		if err == nil {
			return conn, nil
		}
		time.Sleep(50 * time.Millisecond)
	}
	return nil, fmt.Errorf("mpv IPC socket %s never appeared after 2.5s", ipcSocketPath)
}
