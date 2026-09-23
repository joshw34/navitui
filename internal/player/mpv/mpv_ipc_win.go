//go:build windows

package mpv

import (
	"fmt"
	"net"
	"time"

	"github.com/Microsoft/go-winio"
)

const ipcSocketPath = `\\.\pipe\navitui-mpv`

func connectIPC() (net.Conn, error) {
	for range 50 {
		timeout := 50 * time.Millisecond
		conn, err := winio.DialPipe(ipcSocketPath, &timeout)
		if err == nil {
			return conn, nil
		}
		time.Sleep(50 * time.Millisecond)
	}
	return nil, fmt.Errorf("mpv IPC pipe %s never appeared after 2.5s", ipcSocketPath)
}
