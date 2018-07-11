package main

import (
	"errors"
	"flag"
	"os"
	"time"

	"github.com/jaimelopez/chihuahua/executor"
	"github.com/jaimelopez/chihuahua/logger"
	"github.com/jaimelopez/chihuahua/printer"
	"github.com/jaimelopez/chihuahua/storage"
)

const appHeader = `
   ___ _     _ _                 _                 
  / __\ |__ (_) |__  _   _  __ _| |__  _   _  __ _ 
 / /  | '_ \| | '_ \| | | |/ _. | '_ \| | | |/ _. |
/ /___| | | | | | | | |_| | (_| | | | | |_| | (_| |
\____/|_| |_|_|_| |_|\____|\____|_| |_|\____|\____|
 `

func main() {
	group := flag.String("group", "", "Group name of metrics to store *")
	driver := flag.String("storage", "", "Driver to store results (elastic|file) *")
	destination := flag.String("destination", "", "Storage destination *")
	duration := flag.Duration("duration", 1*time.Second, "Time to execute benchmarks")
	threshold := flag.Uint64("threshold", 15, "Threshold percent to determine performance is good enough")
	save := flag.Bool("save", false, "Results will be saved if results are higher than previous")
	force := flag.Bool("force", false, "Force to save results even if they are worse")
	results := flag.Bool("results", false, "Shows results as table")
	debug := flag.Bool("debug", false, "Shows traces for debugging")
	flag.Parse()

	if *group == "" || *driver == "" || *destination == "" {
		usage()
	}

	strg, err := storage.New(*group, *driver, *destination)
	if err != nil {
		logger.Error("setting up storage", err)
	}

	latestResult, err := strg.GetLatest()
	if err != nil {
		logger.Error("retrieving latest benchmarks results", err)
	}

	currentResult, err := executor.Run(*duration, *debug)
	if err != nil {
		logger.Error("executing benchmarks", err)
	} else if len(*currentResult) == 0 {
		logger.Error("executing benchmarks", errors.New("no test found"))
	}

	succeed, comparision := executor.Compare(latestResult, currentResult, uint(*threshold))

	if *results {
		printer.Print(comparision)
	}

	if *save && (succeed || *force) {
		err = strg.Persist(currentResult)
		if err != nil {
			logger.Error("saving results:", err)
		}
	}

	if succeed {
		exit("Good perfomance dude!", 0)
	} else if *force {
		exit("Bad performance but forcing...", 0)
	}

	exit("Bad performance!", -1)
}

func usage() {
	logger.Info(appHeader)
	flag.PrintDefaults()

	exit("", 0)
}

func exit(msg string, status int) {
	logger.Info(msg)
	os.Exit(status)
}
