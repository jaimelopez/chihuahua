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
	equalSymbol  string = "➔"
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
			if metric.CurrentValue == metric.LatestValue {
				worse = equalSymbol
			} else if metric.IsWorse() {
				worse = worseSymbol
			}

			fmt.Printf("%2s %4d%% %s (%d/%d)\n",
				worse,
				int(metric.Diff),
				metric.Metric,
				int(metric.CurrentValue),
				int(metric.LatestValue))
		}

		fmt.Println()
	}
}
