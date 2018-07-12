package printer

import (
	"os"
	"strconv"

	"github.com/jaimelopez/chihuahua/executor"
	"github.com/olekukonko/tablewriter"
)

const (
	checkSymbol string = "✅"
	crossSymbol string = "❌"
)

// Print comparision against standard output
func Print(comp []executor.Comparision) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Valid", "Worse", "Test", "Metric", "Current", "Latest", "Diff"})
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")

	for _, comp := range comp {
		for _, metric := range comp.Metrics {
			worse := checkSymbol
			if metric.IsWorse() {
				worse = crossSymbol
			}

			valid := checkSymbol
			if !metric.Valid {
				valid = crossSymbol
			}

			table.Append([]string{
				valid,
				worse,
				comp.Test,
				metric.Metric,
				strconv.FormatFloat(metric.CurrentValue, 'f', 0, 64),
				strconv.FormatFloat(metric.LatestValue, 'f', 0, 64),
				strconv.FormatFloat(metric.Diff, 'f', 0, 64),
			})
		}
	}

	table.Render()
}
