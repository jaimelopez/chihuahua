package logger

import (
	"log"
	"os"
)

var l = log.New(os.Stderr, "", 0)

// Info prints an info message directly to standard output
func Info(v ...interface{}) {
	l.Println(v...)
}

// Error generates an error trace into standard output and ends the application
func Error(v ...interface{}) {
	l.Fatalln(append([]interface{}{"ERROR:"}, v...)...)
}
