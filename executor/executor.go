package executor

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"golang.org/x/tools/benchmark/parse"
)

// Run benchmarks
func Run(duration time.Duration, debug bool) (*Result, error) {
	var output bytes.Buffer

	params := []string{
		"test",                           // Execute tests
		"-p 1",                           // Parallelization between different packages;
		"-parallel 1",                    // Number of parallel test running for same package
		"-count 1",                       // Avoids to cache results
		"-cpu 2",                         // Number of cpus (using always 2 to avoid missing results)
		"-bench .",                       // Run all benchs
		"-run NONE",                      // Run only those tests and examples matching the regular expression
		"-benchtime" + duration.String(), // Tests duration
		"-benchmem",                      // Includes memmory into report
		"./...",
	}

	cmd := exec.Command("go", params...)
	cmd.Stdout = io.Writer(&output)
	cmd.Stderr = io.Writer(os.Stderr)

	if debug {
		cmd.Stdout = io.MultiWriter(os.Stdout, &output)
	}

	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	return Parse(&output)
}

// FromFile takes benchmarks results directly printed to from file
func FromFile(file string) (*Result, error) {
	c, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	return Parse(bytes.NewBuffer(c))
}

// Parse parse benchmarks results
func Parse(output *bytes.Buffer) (*Result, error) {
	buf := filter(output)

	set, err := parse.ParseSet(buf)
	if err != nil {
		return nil, err
	}

	result := &Result{}

	for name, values := range set {
		(*result)[name] = NewTestResult(values[0])
	}

	return result, nil
}

func filter(output *bytes.Buffer) *bytes.Buffer {
	buf := &bytes.Buffer{}
	scan := bufio.NewScanner(output)

	for scan.Scan() {
		t := scan.Text()
		fields := strings.Fields(t)

		if strings.HasPrefix(fields[0], "Benchmark") {
			_, _ = buf.WriteString("\n" + fields[0])
		}

		for idx, val := range fields {
			if _, err := strconv.Atoi(val); err == nil {
				_, _ = buf.WriteString(" " + val)

				if strings.Contains(fields[idx+1], "/") {
					_, _ = buf.WriteString(" " + fields[idx+1])
				}
			}
		}
	}

	return buf
}
