package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cmonge21/gator/internal/config"
)

func main() {
	// Read the config file
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	// Create a new state with the config
	state := &State{
		Config: cfg,
	}

	// Initialize the commands map
	cmds := &Commands{
		Handlers: make(map[string]func(*State, Command) error),
	}

	// Register the login handler
	cmds.Register("login", handlerLogin)

	// Check if we have enough command line arguments
	if len(os.Args) < 2 {
		fmt.Println("Error: not enough arguments")
		os.Exit(1)
	}

	// Parse the command name and arguments
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	// Create a command
	cmd := Command{
		Name: cmdName,
		Args: cmdArgs,
	}

	// Run the command
	err = cmds.Run(state, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
