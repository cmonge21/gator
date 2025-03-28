package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cmonge21/gator/internal/config"
)

type State struct {
	Config *config.Config
}

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	Handlers map[string]func(*State, Command) error
}

func handlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("username is required")
	}

	if err := s.Config.SetUser(cmd.Args[0]); err != nil {
		return err
	}

	fmt.Printf("User is set to %s\n", cmd.Args[0])
	return nil
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.Handlers[name] = f
}

func (c *Commands) Run(s *State, cmd Command) error {
	handler, exists := c.Handlers[cmd.Name]
	if !exists {
		return fmt.Errorf("unknwown command: %s", cmd.Name)
	}

	return handler(s, cmd)
}

func main() {
	// Read the config file
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	// Create a new state with the config
	state := &State{
		Config: &cfg,
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
