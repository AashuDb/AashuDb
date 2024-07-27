package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func performMetaAction(command string) {
	switch command {
	case ".exit":
		// Returning 0 as an exit code means EXIT_SUCCESS.
		os.Exit(0)
	case ".help":
		fmt.Println("Not Yet Implemented")
	default:
		fmt.Println("Unrecognized meta command:", command)
		fmt.Println("For more info use command, .help;")
	}
}

func main() {

	// Initilaizes an infinite loop for userQuery and reads Input Buffer
	// till a semicolon(;) is encountered.
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("~~~~~~~")
		fmt.Print("aashudb > ")

		userQuery, _ := reader.ReadString(';')
		userQuery = strings.TrimSpace(strings.TrimSuffix(userQuery, ";"))

		// Any query with a initial symbol dot(.) is defined as a META COMMAND
		// Checks for META COMMAND and performs required action, otherwise,
		// treats as a SQL Query.
		if userQuery[0] == '.' {
			performMetaAction(userQuery)
		} else {
			fmt.Println("")
			fmt.Println(userQuery)
		}
	}

}
