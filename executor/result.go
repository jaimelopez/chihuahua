package executor

import "golang.org/x/tools/benchmark/parse"

// Result todo
type Result map[string]*TestResult

// TestResult todo
type TestResult parse.Benchmark

// NewTestResult todo
func NewTestResult(b *parse.Benchmark) *TestResult {
	tr := TestResult(*b)

	return &tr
}

// Rate todo
func (tr *TestResult) Rate() uint64 {
	vns := tr.NsPerOp * 100
	vabo := int(tr.AllocedBytesPerOp) * -10
	vapo := int(tr.AllocsPerOp) * -5
	vt := (uint64(vns) + uint64(vabo) + uint64(vapo)) / 10000

	if vt < 1 {
		return 1
	}

	return vt
}
