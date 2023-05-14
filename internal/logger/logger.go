package logger

import (
	"log"
	"os"
)

const (
	stdFlags = log.LstdFlags | log.LUTC
)

var (
	Info = log.New(os.Stdout, "INFO:\t", stdFlags)
	Warn = log.New(os.Stderr, "WARN:\t", stdFlags|log.Lshortfile)
	Err  = log.New(os.Stderr, "ERR: \t", stdFlags|log.Lshortfile)

	infoFile *os.File
	warnFile *os.File
	errFile  *os.File
)

func Init(infoPath string, warnPath string, errPath string) {
	openLog := func(p string) (*os.File, error) {
		return os.OpenFile(p, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	}
	if infoFile, err := openLog(infoPath); err == nil {
		Info.SetOutput(infoFile)
	}
	if warnFile, err := openLog(warnPath); err == nil {
		Warn.SetOutput(warnFile)
	}
	if errFile, err := openLog(errPath); err == nil {
		Err.SetOutput(errFile)
	}
}

func Close() {
	if infoFile != nil {
		infoFile.Close()
	}
	if warnFile != nil {
		warnFile.Close()
	}
	if errFile != nil {
		errFile.Close()
	}
}
