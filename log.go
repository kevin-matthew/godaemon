package daemon

import (
	"fmt"
	"io"
	"os"
)


const (
	red   = "\033[1;31m"
	yellow= "\033[1;33m"
	white = "[\033[1;37m"
	cyan  = "[\033[1;36m"
)

/*
// TODO:  implement the below
// all levels from /usr/include/sys/syslog.h are included here.
type Level int
const (
	LEVEL_EMERG Level = iota
	LEVEL_ALERT
	LEVEL_CRIT
	LEVEL_ERR
	LEVEL_WARNING
	LEVEL_NOTICE
	LEVEL_INFO
	LEVEL_DEBUG
)

func (Level s) ToString() string {
	// TODO
}


*/

// Feel free to set this to whatever you'd like. All output of a daemon should be going to the same place.
// Thus we don't need to use os.Stderr for example (we should only use that if there was a
// problem starting the daemon).
var Output io.Writer = os.Stdout

func logColor(prefix string, color string, message string) {
	_,_ = fmt.Fprintf(Output, "[ %s%.6s \033[0m] %s\n", color,prefix,message)
}

func Emerg(message string) {
	logColor("Emerg", red, message)
}

func Error(message string) {
	logColor("Error", red, message)
}

func Crit(message string) {
	logColor("Crit", red, message)
}

func Alert(message string) {
	logColor("Alert", yellow, message)
}

func Warning(message string) {
	logColor("Warning", yellow, message)
}

func Notice(message string) {
	logColor("Notice", yellow, message)
}

func Info(message string) {
	logColor("Alert", white, message)
}

func Debug(message string) {
	logColor("Debug", cyan, message)
}


