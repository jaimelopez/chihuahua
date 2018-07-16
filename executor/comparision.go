package executor

// These constants specifies the display name for some metrics
const (
	MetricNsPerOpDisplay           string = "time"        // NsPerOp metric
	MetricAllocedBytesPerOpDisplay string = "memory"      // AllocedBytesPerOp metric
	MetricAllocsPerOpDisplay       string = "allocations" // AllocsPerOp metric
)

// Measures used in the metrics
const (
	MetricNsPerOpMeasure           string = "ns"    // Nanoseconds per operation
	MetricAllocedBytesPerOpMeasure string = "bytes" // Bytes allocated per operation
	MetricAllocsPerOpMeasure       string = "times" // Allocation times per operation
)

// Comparision represents a whole comparision between two results
type Comparision struct {
	Name    string
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
