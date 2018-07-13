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
				calculate(MetricNsPerOpDisplay, currentBench.NsPerOp, latestBench.NsPerOp, threshold, MetricNsPerOpMeasure),
				calculate(MetricAllocedBytesPerOpDisplay, float64(currentBench.AllocedBytesPerOp), float64(latestBench.AllocedBytesPerOp), threshold, MetricAllocedBytesPerOpMeasure),
				calculate(MetricAllocsPerOpDisplay, float64(currentBench.AllocsPerOp), float64(latestBench.AllocsPerOp), threshold, MetricAllocsPerOpMeasure),
			},
		}

		list = append(list, cmp)

		if !cmp.IsValid() {
			succeed = false
		}
	}

	return succeed, list
}

func calculate(metric string, current float64, latest float64, threshold uint, measure string) MetricComparision {
	diff := (float64(latest)/float64(current))*100 - 100

	if latest == 0 {
		diff = 0
	}

	valid := diff >= 0 || math.Abs(diff) <= float64(threshold)

	return MetricComparision{
		Metric:       metric,
		CurrentValue: current,
		LatestValue:  latest,
		Measure:      measure,
		Diff:         diff,
		Valid:        valid,
	}
}
