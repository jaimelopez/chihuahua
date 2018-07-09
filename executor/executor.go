package executor

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"time"

	"golang.org/x/tools/benchmark/parse"
)

// Execute todo
func Execute(duration time.Duration) (*Result, error) {
	var output bytes.Buffer

	cmd := exec.Command("go", "test", "-bench", ".", "-run", "NONE", "-benchtime", duration.String(), "-benchmem")
	cmd.Stdout = io.Writer(&output)
	cmd.Stderr = io.Writer(os.Stderr)

	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	results, err := parse.ParseSet(&output)
	if err != nil {
		return nil, err
	}
	return &Result{
		ID:    "",
		Tests: results,
	}, nil
}
