package printer

import (
	"fmt"

	"github.com/jaimelopez/chihuahua/executor"
)

const (
	checkSymbol  string = "✅"
	crossSymbol  string = "❌"
	betterSymbol string = "▲"
	worseSymbol  string = "▼"
)

// Print comparision against standard output
func Print(comparisions []executor.Comparision) {
	for _, comp := range comparisions {
		valid := crossSymbol
		if comp.IsValid() {
			valid = checkSymbol
		}

		fmt.Println(valid, comp.Test)

		for _, metric := range comp.Metrics {
			worse := betterSymbol
			if metric.IsWorse() {
				worse = worseSymbol
			}

			diff := fmt.Sprintf("(%d/%d) %d%%", int(metric.CurrentValue), int(metric.LatestValue), int(metric.Diff))

			fmt.Println(" ", worse, metric.Metric, diff)
		}

		fmt.Println()
	}
}
