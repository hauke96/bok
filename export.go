package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/hauke96/sigolo"
)

func exportToCSV(store *Store, file string) {
	var buf bytes.Buffer

	for _, e := range store.Entries {
		buf.WriteString(fmt.Sprintf("\"%d,%d\",", e.Amount/100, e.Amount%100))
		buf.WriteString(fmt.Sprintf("\"%s\",", e.Note))
		buf.WriteString(fmt.Sprintf("\"%s\",", e.Date.Format("02.01.2006")))
		buf.WriteString(fmt.Sprintf("\"%s\"", e.Category))
		buf.WriteString("\n")
	}

	err := ioutil.WriteFile(file, buf.Bytes(), 0644)
	sigolo.FatalCheck(err)
}
