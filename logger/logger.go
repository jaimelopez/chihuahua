package logger

import (
	"fmt"
	"os"
)

// Info todo
func Info(v ...interface{}) {
	fmt.Println(v...)
}

// Error todo
func Error(where string, err error) {
	fmt.Println("[ERROR]", where, ":", err)

	os.Exit(-1)
}
