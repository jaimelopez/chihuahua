package parameters

import (
	"flag"
	"time"
)

// Params defines behaviour to run the app
type Parameters struct {
	Group       *string
	Driver      *string
	Destination *string
	FromFile    *string
	Duration    *time.Duration
	Threshold   *uint64
	Save        *bool
	Force       *bool
	Results     *bool
	Debug       *bool
}

// NewFromFlags instantiate a new params struct from flags
func NewFromFlags() *Parameters {
	params := &Parameters{
		Group:       flag.String("group", "", "Group name of metrics to store *"),
		Driver:      flag.String("storage", "", "Driver to store results (elastic|file) *"),
		Destination: flag.String("destination", "", "Storage destination *"),
		FromFile:    flag.String("fromfile", "", "Takes results to analyze from file instead running benchmarks"),
		Duration:    flag.Duration("duration", 1*time.Second, "Time to execute benchmarks"),
		Threshold:   flag.Uint64("threshold", 15, "Threshold percent to determine performance is good enough"),
		Save:        flag.Bool("save", false, "Results will be saved if results are higher than previous"),
		Force:       flag.Bool("force", false, "Forces to save results even if they are worse"),
		Results:     flag.Bool("results", false, "Shows results as table"),
		Debug:       flag.Bool("debug", false, "Shows traces for debugging"),
	}

	flag.Parse()

	return params
}
