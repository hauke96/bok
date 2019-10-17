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
