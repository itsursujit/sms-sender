package log

import (
	"io/ioutil"
	"log"
	"os"
)

var (
	ExitOnFatal = true
	Trace       = log.New(ioutil.Discard,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warn = log.New(os.Stdout,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(os.Stderr,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
)

func Fatal(v ...interface{}) {
	Error.Print(v...)
	if ExitOnFatal {
		os.Exit(1)
	}
}

func WarnIfErr(description string, err error) {
	if err != nil {
		Warn.Println(description, err)
	}
}

func ErrIfErr(description string, err error) {
	if err != nil {
		Error.Println(description, err)
	}
}
