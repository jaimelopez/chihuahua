package executor

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"time"

	"golang.org/x/tools/benchmark/parse"
)

// Run benchmarks
func Run(duration time.Duration, debug bool) (*Result, error) {
	var output bytes.Buffer

	cmd := exec.Command("go", "test", "-bench", ".", "-run", "NONE", "-benchtime", duration.String(), "-benchmem")
	cmd.Stdout = io.Writer(&output)
	cmd.Stderr = io.Writer(os.Stderr)

	if debug {
		cmd.Stdout = io.MultiWriter(os.Stdout, &output)
	}

	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	return Parse(output)
}

// Parse parse benchmarks results
func Parse(output bytes.Buffer) (*Result, error) {
	set, err := parse.ParseSet(&output)
	if err != nil {
		return nil, err
	}

	result := &Result{}

	for name, values := range set {
		(*result)[name] = NewTestResult(values[0])
	}

	return result, nil
}
