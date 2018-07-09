package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/jaimelopez/chihuahua/executor"
)

const appHeader = `
   ___ _     _ _                 _                 
  / __\ |__ (_) |__  _   _  __ _| |__  _   _  __ _ 
 / /  | '_ \| | '_ \| | | |/ _. | '_ \| | | |/ _. |
/ /___| | | | | | | | |_| | (_| | | | | |_| | (_| |
\____/|_| |_|_|_| |_|\____|\____|_| |_|\____|\____|

 Making Sergey-bus great again
 `

// const usage = `Usage: chihuahua [options...]
// Options:
// 	-storage    	driver to store results
// 	-connection		storage connection string
// 	-duration		time to execute benchmarks
// 	-threshold		percent of threshold to determine performance is good enough
// 	-save  			results will be saved if results are higher than previous
// 	-force			force to save results even if they are worse`

func main() {
	storage := flag.String("storage", "", "Driver to store results (elastic|file}) *")
	destination := flag.String("destination", "", "Storage destination *")
	duration := flag.Duration("duration", 1*time.Second, "Time to execute benchmarks")
	threshold := flag.Uint64("threshold", 15, "Threshold percent to determine performance is good enough")
	save := flag.Bool("save", false, "Results will be saved if results are higher than previous")
	force := flag.Bool("force", false, "Force to save results even if they are worse")
	flag.Parse()

	if *storage == "" || *destination == "" {
		usage()
		return
	}

	result, err := executor.Execute(*duration)

	if err != nil {
		fmt.Println("ERROR executing benchmarks:", err)
		return
	}

	fmt.Println(result)
	fmt.Println(duration, threshold, save, force)
}

func usage() {
	fmt.Println(appHeader)
	flag.PrintDefaults()
	fmt.Println()
}
