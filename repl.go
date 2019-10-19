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

func RunRepl(store *Store) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("> ")
	for scanner.Scan() {
		input := scanner.Text()

		for i := 0; i < len(input); i++ {
			cmd := input[i]

			// commands can be forced by typing "!" after the command
			force := false
			if i+1 < len(input) && input[i+1] == '!' {
				force = true
				i++
			}

			switch cmd {
			case 'q':
				if !store.Dirty || store.Dirty && force {
					sigolo.Info("Bye")
					goto ReplEnd
				} else {
					sigolo.Info("There are unsaved changes, please save wit 'w' before exiting.")
				}
			case 'a':
				err := replAddEntry(scanner, store)
				if err != nil {
					sigolo.Info("Error adding entry: " + err.Error())
				}
			case 'w':
				store.SaveStore()
				sigolo.Info("Saved")
			default:
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

func replAddEntry(scanner *bufio.Scanner, store *Store) error {
	var dateString string
	var amountString string
	var description string
	var category string

	sigolo.Info("Add new entry:")

	fmt.Print("  Date: ")
	scanner.Scan()
	dateString = scanner.Text()

	fmt.Print("  Amount: ")
	scanner.Scan()
	amountString = scanner.Text()

	fmt.Print("  Description: ")
	scanner.Scan()
	description = scanner.Text()

	fmt.Print("  Category: ")
	scanner.Scan()
	category = scanner.Text()

	// Convert strings to right type

	amountString = strings.ReplaceAll(amountString, ",", ".")
	amountFloat, err := strconv.ParseFloat(amountString, 64)
	if err != nil {
		return err
	}

	amount := int(amountFloat * 100.0)

	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return err
	}

	store.AddEntry(amount, date, description, category)

	return nil
}
