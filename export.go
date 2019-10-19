package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/hauke96/sigolo"
)

func export(store *Store, format string, fileWithoutExtension string) error {
	var err error
	var file string

	switch format {
	case "csv":
		file = fileWithoutExtension + ".csv"
		err = exportToCSV(store, file)
	default:
		err = errors.New(fmt.Sprintf("Unsupported format '%s'", format))
	}

	if err != nil {
		return err
	}

	sigolo.Info("Exporting finished succesfully")
	sigolo.Info("  Format: %s", format)
	sigolo.Info("  File: %s", file)

	return nil
}

func exportToCSV(store *Store, file string) error {
	var buf bytes.Buffer

	for _, e := range store.Entries {
		buf.WriteString(fmt.Sprintf("\"%d,%d\",", e.Amount/100, e.Amount%100))
		buf.WriteString(fmt.Sprintf("\"%s\",", e.Note))
		buf.WriteString(fmt.Sprintf("\"%s\",", e.Date.Format("02.01.2006")))
		buf.WriteString(fmt.Sprintf("\"%s\"", e.Category))
		buf.WriteString("\n")
	}

	err := ioutil.WriteFile(file, buf.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}
