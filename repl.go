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
	var cmd string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("> ")
	for scanner.Scan() {
		cmd = scanner.Text()

		if cmd == "q" || cmd == "quit" || cmd == "exit" {
			break
		}
		switch cmd {
		case "add", "a":
			err := replAddEntry(scanner, store)
			if err != nil {
				sigolo.Info("Error adding entry: " + err.Error())
			}
		case "write", "w":
			store.SaveStore()
		default:
			sigolo.Info("Unknown command '%s'", cmd)
		}

		fmt.Print("> ")
	}

	if err := scanner.Err(); err != nil {
		sigolo.Fatal("Error reading standard input: ", err.Error())
	}
}

func replAddEntry(scanner *bufio.Scanner, store *Store) error {
	var dateString string
	var amountString string
	var description string

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

	store.AddEntry(amount, description, date)

	return nil
}
