package executor

import "golang.org/x/tools/benchmark/parse"

// Result todo
type Result struct {
	ID    string
	Tests map[string][]*parse.Benchmark
}
