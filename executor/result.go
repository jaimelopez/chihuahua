package executor

import "golang.org/x/tools/benchmark/parse"

// Result of a complete benchmark execution
type Result map[string]*TestResult

// TestResult represents an specific test benchmark result
type TestResult parse.Benchmark

// NewTestResult generates a TestResult from a benchmark parse
func NewTestResult(b *parse.Benchmark) *TestResult {
	tr := TestResult(*b)

	return &tr
}
