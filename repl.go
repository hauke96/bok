package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hauke96/sigolo"
)

// RunRepl starts the read-eval-print loop
func RunRepl(store *Store) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("> ")
	for scanner.Scan() {
		input := scanner.Text()

		// When the user enters more than one character
		for i := 0; i < len(input); i++ {
			cmd := input[i]

			// commands can be forced by typing "!" after the command
			force := false
			if i+1 < len(input) && input[i+1] == '!' {
				force = true
				i++
			}

			// Evalualte the command
			switch cmd {
			case 'q': // -> quit
				if !store.Dirty || store.Dirty && force {
					sigolo.Info("Bye")
					goto ReplEnd // out of both loops
				} else {
					sigolo.Info("There are unsaved changes, please save wit 'w' before exiting.")
				}
			case 'a': // -> add
				err := replAddEntry(scanner, store)
				if err != nil {
					sigolo.Info("Error adding entry: " + err.Error())
				}
			case 'w': // -> write to disk
				store.SaveStore()
				sigolo.Info("Saved")
			default: // -> unknown command
				sigolo.Info("Unknown command '%c' at pos %d", cmd, i)
			}
		}

		fmt.Print("> ")
	}

ReplEnd:

	if err := scanner.Err(); err != nil {
		sigolo.Fatal("Error reading standard input: ", err.Error())
	}
}

// replAddEntry starts is own little REPL by asking the user for the necessary
// data to add an accounting entry.
func replAddEntry(scanner *bufio.Scanner, store *Store) error {
	var dateString string
	var amountString string
	var description string
	var category string
	var lastEntry *Entry
	if len(store.Entries) != 0 {
		lastEntry = store.Entries[len(store.Entries)-1]
	}

	sigolo.Info("Add new entry:")

	// TODO Maybe reuse a lot of code here

	dateString = askForData(scanner, "Date", lastEntry.Date.Format("2006-01-02"))

	amountString = askForData(scanner, "Amounnt", fmt.Sprintf("%d.%d", lastEntry.Amount/100, lastEntry.Amount%100))

	description = askForData(scanner, "Description", lastEntry.Note)

	category = askForData(scanner, "Category", lastEntry.Category)

	// Convert strings to right type
	amountString = strings.ReplaceAll(amountString, ",", ".")
	amountFloat, err := strconv.ParseFloat(amountString, 64)
	if err != nil {
		return err
	}
	amount := int(amountFloat * 100.0)

	// Currently only the "yyyy-mm-dd" format is supported
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return err
	}

	store.AddEntry(amount, date, description, category)

	return nil
}

// askForData scans for input from the user. If no input is given, the fallback value is returned
func askForData(scanner *bufio.Scanner, text, fallback string) string {
	fmt.Printf("  %s (%s): ", text, fallback)

	scanner.Scan()
	input := scanner.Text()

	if len(fallback) != 0 && len(input) == 0 {
		input = fallback
		sigolo.Info("    (%s)", input)
	}

	return input
}
