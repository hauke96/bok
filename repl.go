package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/hauke96/sigolo"
)

func RunRepl() {
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
			replAddEntry(scanner)
		default:
			sigolo.Info("Unknown command '%s'", cmd)
		}

		fmt.Print("> ")
	}

	if err := scanner.Err(); err != nil {
		sigolo.Fatal("Error reading standard input: ", err.Error())
	}
}

func replAddEntry(scanner *bufio.Scanner) {
	var date string
	var amount string
	var description string

	sigolo.Info("Add new entry:")

	fmt.Print("  Date: ")
	scanner.Scan()
	date = scanner.Text()

	fmt.Print("  Amount: ")
	scanner.Scan()
	amount = scanner.Text()

	fmt.Print("  Description: ")
	scanner.Scan()
	description = scanner.Text()

	// TODO addEntry(date, amount, description)

	sigolo.Debug("Added entry [%s, %s, '%s']", date, amount, description)
}
