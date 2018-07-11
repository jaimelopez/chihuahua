package logger

import (
	"fmt"
	"os"
)

// Info prints an info message directly to standard output
func Info(v ...interface{}) {
	fmt.Println(v...)
}

// Error generates an error trace into standard output and ends the application
func Error(where string, err error) {
	fmt.Println("[ERROR]", where, ":", err)

	os.Exit(-1)
}
