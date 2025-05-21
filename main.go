package main

import (
	"fmt"
	"io/ioutil"
	"nobl/repl"
	"nobl/web"
	"os"
	"os/user"
	"strings"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	// Check command line arguments
	if len(os.Args) > 1 {
		// Check for web server flag
		if os.Args[1] == "--web" || os.Args[1] == "-w" {
			port := 8080
			if len(os.Args) > 2 {
				fmt.Sscanf(os.Args[2], "%d", &port)
			}
			web.ServeWebREPL(port)
			return
		}

		// Get the file path from command line arguments
		filePath := os.Args[1]

		// Read the file
		fileBytes, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading file: %s\n", err)
			os.Exit(1)
		}

		// Convert file content to a reader
		fileContent := string(fileBytes)
		reader := strings.NewReader(fileContent)

		// Execute the file content
		repl.Start(reader, os.Stdout)
	} else {
		// Start interactive REPL if no file provided
		fmt.Printf("Hello %s! This is the Nobl programming language!\n", user.Username)
		fmt.Printf("Feel free to type in commands\n")
		repl.Start(os.Stdin, os.Stdout)
	}
}
