package executor

const (
	// MetricNsPerOpDisplay specifies name to display for NsPerOp metric
	MetricNsPerOpDisplay string = "time"
	// MetricAllocedBytesPerOpDisplay specifies name to display for AllocedBytesPerOp metric
	MetricAllocedBytesPerOpDisplay string = "memory"
	// MetricAllocsPerOpDisplay specifies name to display for AllocsPerOp metric
	MetricAllocsPerOpDisplay string = "allocations"
)

const (
	// MetricNsPerOpMeasure specifies used measure in NsPerOp metric
	MetricNsPerOpMeasure string = "ns"
	// MetricAllocedBytesPerOpMeasure specifies used measure in AllocedBytesPerOp metric
	MetricAllocedBytesPerOpMeasure string = "bytes"
	// MetricAllocsPerOpMeasure specifies used measure in AllocsPerOp metric
	MetricAllocsPerOpMeasure string = "times"
)

// Comparision represents a whole comparision between two results
type Comparision struct {
	Test    string
	Metrics []MetricComparision
}

// IsValid determines if is valid even being worse but stills over threshold
func (cmp *Comparision) IsValid() bool {
	for _, res := range cmp.Metrics {
		if !res.Valid {
			return false
		}
	}

	return true
}

// IsWorse indicates if comparasion is fully or partially worse than previous one
func (cmp *Comparision) IsWorse() bool {
	for _, res := range cmp.Metrics {
		if res.IsWorse() {
			return true
		}
	}

	return false
}

// MetricComparision represents a metric comparision against latest results
type MetricComparision struct {
	Metric       string
	CurrentValue float64
	LatestValue  float64
	Measure      string
	Diff         float64
	Valid        bool
}

// IsWorse defines whether metric is worse than previous one or not
func (mc *MetricComparision) IsWorse() bool {
	return mc.Diff < 0
}
