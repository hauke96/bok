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
	appFile  = app.Flag("file", "The accounting file to use").Default("account.json").String()

	appExport       = app.Command("export", "Exports the store to a different format")
	appExportFormat = appExport.Flag("format", "This is the format the accounting data gets converted to. Supported formats: csv").Default("csv").String()
	appExportOutput = appExport.Flag("output", "The output file (without extension)").Default("account_export").String()
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
	}
}

func main() {
	configureCliArgs()

	cmd, err := app.Parse(os.Args[1:])
	sigolo.FatalCheck(err)

	configureLogging()

	store := ReadStore(*appFile)

	sigolo.Info("Welcome to bokhald")
	sigolo.Debug("Read store: %v", store)

	switch cmd {
	case appExport.FullCommand():
		exportToCSV(store, *appExportOutput+".csv")
	default:
		RunRepl(store)
	}
}
