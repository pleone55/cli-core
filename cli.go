package cli

import (
	"flag"
	"fmt"
	"os"
)

// CLI represents the command-line interface application
type CLI struct {
	Name        string
	Description string
	Commands    []Command
	GlobalFlags *flag.FlagSet
}

// NewCLI creates a new CLI instance
func NewCLI(name, description string) *CLI {
	return &CLI{
		Name:        name,
		Description: description,
		Commands:    []Command{},
		GlobalFlags: flag.NewFlagSet(name, flag.ExitOnError),
	}
}

// AddCommand adds a command to the CLI
func (cli *CLI) AddCommand(command Command) {
	cli.Commands = append(cli.Commands, command)
}

// SetGlobalFlag sets a global flag for the CLI
func (cli *CLI) SetGlobalFlag(name string, value string, usage string) {
	cli.GlobalFlags.String(name, value, usage)
}

// Usage prints the usage information for the CLI
func (cli *CLI) Usage() {
	fmt.Fprintf(os.Stderr, "Usage: of %s [command]\n\n", cli.Name)
	fmt.Fprintf(os.Stderr, "%s\n\n", cli.Description)
	fmt.Fprintf(os.Stderr, "Available commands:\n")
	for _, cmd := range cli.Commands {
		fmt.Fprintf(os.Stderr, "  %s: %s\n", cmd.Name, cmd.Description)
	}
	fmt.Fprintf(os.Stderr, "Usage %s help [command] for more information on a specific command.\n", cli.Name)
}

// Run executes the CLI application
func (cli *CLI) Run(args []string) error {
	if len(args) < 1 {
		cli.Usage()
		return nil
	}

	// Parse global flags
	if err := cli.GlobalFlags.Parse(args); err != nil {
		return err
	}

	// Find the command to execute
	commandName := args[0]
	for _, cmd := range cli.Commands {
		if cmd.Name == commandName {
			// Parse command-specific flags
			cmdFlags := flag.NewFlagSet(cmd.Name, flag.ExitOnError)
			for _, f := range cmd.Flags {
				cmdFlags.Var(f.Value, f.Name, f.Usage)
			}
			if err := cmdFlags.Parse(args[1:]); err != nil {
				return err
			}

			// Run the command
			return cmd.Run(cmdFlags.Args())
		}
	}
	// If command not found, print usage
	fmt.Fprintf(os.Stderr, "Unknown command: %s\n", commandName)
	cli.Usage()
	return fmt.Errorf("unknown command: %s", commandName)
}
