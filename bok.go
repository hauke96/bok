package main

import (
	"os"

	"github.com/hauke96/kingpin"
	"github.com/hauke96/sigolo"
)

const VERSION string = "v0.0.1"

var (
	app      = kingpin.New("bokhald", "A simple accounting tool")
	appDebug = app.Flag("verbose", "Verbose mode, showing additional debug information").Short('v').Bool()

	appAdd         = app.Command("add", "Adds a new entry")
	addDate        = appAdd.Arg("date", "The date of the bill, receip or expense. Format: yyyy-mm-dd").Required().String()
	addAmount      = appAdd.Arg("amount", "The amount of money. Format: x,yy").Required().String()
	addDescription = appAdd.Arg("description", "A brief description or key-word").Required().String()
	addFile        = appAdd.Arg("file", "The accounting file to use").Default("account.json").String()

	// TODO remove

	// TODO edit
)

func configureCliArgs() {
	app.Author("Hauke Stieler")
	app.Version(VERSION)

	app.HelpFlag.Short('h').Help("Shows this help message")
}

func configureLogging() {
	sigolo.FormatFunctions[sigolo.LOG_INFO] = sigolo.FormatFunctions[sigolo.LOG_PLAIN]

	if *appDebug {
		sigolo.LogLevel = sigolo.LOG_DEBUG
	} else {
		sigolo.LogLevel = sigolo.LOG_PLAIN
	}
}

func main() {
	configureCliArgs()

	cmd, err := app.Parse(os.Args[1:])
	sigolo.FatalCheck(err)

	configureLogging()

	sigolo.Info("Welcome to bokhald")

	switch cmd {
	case appAdd.FullCommand():
		sigolo.Debug("Use account file '%s'", *addFile)
	}
}
