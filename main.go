package main

import (
	"errors"
	"flag"
	"os"

	"github.com/jaimelopez/chihuahua/executor"
	"github.com/jaimelopez/chihuahua/logger"
	"github.com/jaimelopez/chihuahua/parameters"
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

const (
	successExitCode          int = 0
	badUsageExitCode         int = 0
	performanceErrorExitCode int = 1
)

func main() {
	params := parameters.NewFromFlags()

	if *params.Group == "" || *params.Driver == "" || *params.Destination == "" {
		usage()
	}

	strg, err := storage.New(*params.Group, *params.Driver, *params.Destination)
	if err != nil {
		logger.Error("setting up storage", err)
	}

	latestResult, err := strg.GetLatest()
	if err != nil {
		logger.Error("retrieving latest benchmarks results", err)
	}

	currentResult, err := executor.Run(*params.Duration, *params.Debug)
	if err != nil {
		logger.Error("executing benchmarks", err)
	} else if len(*currentResult) == 0 {
		logger.Error("executing benchmarks", errors.New("no test found"))
	}

	succeed, comparision := executor.Compare(latestResult, currentResult, uint(*params.Threshold))

	if *params.Results {
		printer.Print(comparision)
	}

	if *params.Save && (succeed || *params.Force) {
		err = strg.Persist(currentResult)
		if err != nil {
			logger.Error("saving results:", err)
		}
	}

	if succeed {
		exit("Good perfomance dude!", successExitCode)
	} else if *params.Force {
		exit("Bad performance but forcing...", successExitCode)
	}

	exit("Bad performance!", performanceErrorExitCode)
}

func usage() {
	logger.Info(appHeader)
	flag.PrintDefaults()

	exit("", badUsageExitCode)
}

func exit(msg string, status int) {
	logger.Info(msg)
	os.Exit(status)
}
