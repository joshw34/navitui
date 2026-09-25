package types

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/adrg/xdg"
)

type NavituiLogger struct {
	logFile  *os.File
	file     *log.Logger
	terminal *log.Logger
	both     *log.Logger
	logMutex sync.Mutex
}

func LoggerSetup() *NavituiLogger {
	var l NavituiLogger
	path, err := xdg.DataFile("navitui/navitui.log")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to create logfile: %s\n", err.Error())
		os.Exit(1)
	}

	l.logFile, err = os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to create logfile: %s\n", err.Error())
		os.Exit(1)
	}

	both := io.MultiWriter(os.Stderr, l.logFile)

	l.file = log.New(l.logFile, "", log.Ldate|log.Ltime|log.Lshortfile)
	l.terminal = log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)
	l.both = log.New(both, "", log.Ldate|log.Ltime|log.Lshortfile)

	return &l
}

func (l *NavituiLogger) Close() {
	l.logMutex.Lock()
	defer l.logMutex.Unlock()
	if l.logFile != nil {
		_ = l.logFile.Close()
		l.logFile = nil
		l.file = nil
		l.both = nil
	}
}

func (l *NavituiLogger) File(str string, args ...any) {
	l.output(l.file, str, args...)
}

func (l *NavituiLogger) Terminal(str string, args ...any) {
	l.output(l.terminal, str, args...)
}

func (l *NavituiLogger) Both(str string, args ...any) {
	l.output(l.both, str, args...)
}

func (l *NavituiLogger) output(logger *log.Logger, str string, args ...any) {
	l.logMutex.Lock()
	defer l.logMutex.Unlock()
	if logger == nil {
		return
	}
	s := fmt.Sprintf(str, args...)
	_ = logger.Output(3, s)
}
