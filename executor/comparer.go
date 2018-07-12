package executor

import (
	"math"
)

// Compare two different benchmark results
func Compare(latest *Result, current *Result, threshold uint) (bool, []Comparision) {
	succeed := true
	list := []Comparision{}

	for name, currentBench := range *current {
		latestBench, ok := (*latest)[name]

		if !ok {
			latestBench = &TestResult{}
		}

		cmp := Comparision{
			Test: name,
			Metrics: []MetricComparision{
				calculate("NsPerOp", currentBench.NsPerOp, latestBench.NsPerOp, threshold),
				calculate("AllocedBytesPerOp", float64(currentBench.AllocedBytesPerOp), float64(latestBench.AllocedBytesPerOp), threshold),
				calculate("AllocsPerOp", float64(currentBench.AllocsPerOp), float64(latestBench.AllocsPerOp), threshold),
			},
		}

		list = append(list, cmp)

		if !cmp.IsValid() {
			succeed = false
		}
	}

	return succeed, list
}

func calculate(metric string, current float64, latest float64, threshold uint) MetricComparision {
	diff := (float64(latest)/float64(current))*100 - 100

	if latest == 0 {
		diff = 0
	}

	valid := diff >= 0 || math.Abs(diff) <= float64(threshold)

	return MetricComparision{
		Metric:       metric,
		CurrentValue: current,
		LatestValue:  latest,
		Diff:         diff,
		Valid:        valid,
	}
}
