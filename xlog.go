// Implements the Info, Warn and Error loggers that by default write to os.Stderr
package xlog

import (
	"fmt"
	"io"
	"log"
	"os"
)

type xLoggers struct {
	xinfo       *log.Logger
	xwarn       *log.Logger
	xerror      *log.Logger
	output      io.Writer
	fatalPrefix string
	panicPrefix string
}

var xlog xLoggers

func init() {
	defaultFlags := log.Ldate | log.Lshortfile
	xlog = xLoggers{
		xinfo:       log.New(os.Stderr, "INFO: ", defaultFlags),
		xwarn:       log.New(os.Stderr, "WARN: ", defaultFlags),
		xerror:      log.New(os.Stderr, "ERROR: ", defaultFlags),
		output:      os.Stderr,
		fatalPrefix: "FATAL: ",
		panicPrefix: "PANIC: ",
	}
}

//
// Info print functions
//

func Info(v ...any) {
	xlog.xinfo.Output(2, fmt.Sprint(v...))
}

func Infof(format string, v ...any) {
	xlog.xinfo.Output(2, fmt.Sprintf(format, v...))
}

func Infoln(v ...any) {
	xlog.xinfo.Output(2, fmt.Sprintln(v...))
}

//
// Warn print functions
//

func Warn(v ...any) {
	xlog.xwarn.Output(2, fmt.Sprint(v...))
}

func Warnf(format string, v ...any) {
	xlog.xwarn.Output(2, fmt.Sprintf(format, v...))
}

func Warnln(v ...any) {
	xlog.xwarn.Output(2, fmt.Sprintln(v...))
}

//
// Error print functions
//

func Error(v ...any) {
	xlog.xerror.Output(2, fmt.Sprint(v...))
}

func Errorf(format string, v ...any) {
	xlog.xerror.Output(2, fmt.Sprintf(format, v...))
}

func Errorln(v ...any) {
	xlog.xerror.Output(2, fmt.Sprintln(v...))
}

//
// Map Panic functions to Error
//

func Panic(v ...any) {
	s := fmt.Sprint(v...)
	xlog.xerror.SetPrefix(xlog.panicPrefix)
	xlog.xerror.Output(2, s)
	panic(s)
}

func Panicf(format string, v ...any) {
	s := fmt.Sprintf(format, v...)
	xlog.xerror.SetPrefix(xlog.panicPrefix)
	xlog.xerror.Output(2, s)
	panic(s)
}

func Panicln(v ...any) {
	s := fmt.Sprintln(v...)
	xlog.xerror.SetPrefix(xlog.panicPrefix)
	xlog.xerror.Output(2, s)
	panic(s)
}

//
// Map Fatal functions to Error
//

func Fatal(v ...any) {
	s := fmt.Sprint(v...)
	xlog.xerror.SetPrefix(xlog.fatalPrefix)
	xlog.xerror.Output(2, s)
	os.Exit(1)
}

func Fatalf(format string, v ...any) {
	s := fmt.Sprintf(format, v...)
	xlog.xerror.SetPrefix(xlog.fatalPrefix)
	xlog.xerror.Output(2, s)
	os.Exit(1)
}

func Fatalln(v ...any) {
	s := fmt.Sprintln(v...)
	xlog.xerror.SetPrefix(xlog.fatalPrefix)
	xlog.xerror.Output(2, s)
	os.Exit(1)
}

// Functions that act on all loggers simultaneously
func Enable() {
	xlog.xinfo.SetOutput(xlog.output)
	xlog.xwarn.SetOutput(xlog.output)
	xlog.xerror.SetOutput(xlog.output)
}

func Disable() {
	xlog.xinfo.SetOutput(io.Discard)
	xlog.xwarn.SetOutput(io.Discard)
	xlog.xerror.SetOutput(io.Discard)
}

func SetOutput(w io.Writer) {
	xlog.xinfo.SetOutput(w)
	xlog.xwarn.SetOutput(w)
	xlog.xerror.SetOutput(w)
	xlog.output = w
}

func SetFlags(flag int) {
	xlog.xinfo.SetFlags(flag)
	xlog.xwarn.SetFlags(flag)
	xlog.xerror.SetFlags(flag)
}

func SetPrefixes(infoPrefix, warnPrefix, errPrefix, fatalPrefix, panicPrefix string) {
	xlog.xinfo.SetPrefix(infoPrefix)
	xlog.xwarn.SetPrefix(warnPrefix)
	xlog.xerror.SetPrefix(errPrefix)
	xlog.fatalPrefix = fatalPrefix
	xlog.panicPrefix = panicPrefix
}

func SetShortPrefixes() {
	SetPrefixes("I: ", "W: ", "E: ", "F: ", "P: ")
}
