package executor

// Comparision todo
type Comparision struct {
	Test    string
	Metrics *[]MetricComparision
}

// IsValid todo
func (cmp *Comparision) IsValid() bool {
	for _, res := range *cmp.Metrics {
		if !res.Valid {
			return false
		}
	}

	return true
}

// IsWorse todo
func (cmp *Comparision) IsWorse() bool {
	for _, res := range *cmp.Metrics {
		if res.IsWorse() {
			return true
		}
	}

	return false
}

// MetricComparision todo
type MetricComparision struct {
	Metric       string
	CurrentValue float64
	LatestValue  float64
	Diff         float64
	Valid        bool
}

// IsWorse todo
func (mc *MetricComparision) IsWorse() bool {
	return mc.Diff < 0
}
