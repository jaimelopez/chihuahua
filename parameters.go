package main

import (
	"time"
)

// Parameters defines behaviour to run the app
type Parameters struct {
	Storage    string
	Connection string
	Duration   time.Duration
	Threshold  uint16
	Update     bool
}
