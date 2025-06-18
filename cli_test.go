package cli

import (
	"flag"
	"testing"
)

func TestAddCommand(t *testing.T) {
	cli := NewCLI("testcli", "This is a test CLI application")
	cmd := &Command{
		Name:        "test",
		Description: "This is a test command",
		Flags:       []*flag.Flag{},
		Run: func(args []string) error {
			return nil
		},
	}

	cli.AddCommand(*cmd)

	if len(cli.Commands) != 1 || cli.Commands[0].Name != "test" {
		t.Error("Expected command to be added, but it was not found in the CLI commands")
	}
}

func TestRunCommand(t *testing.T) {
	cli := NewCLI("testcli", "This is a test CLI application")
	cmd := &Command{
		Name:        "test",
		Description: "This is a test command",
		Flags:       []*flag.Flag{},
		Run: func(args []string) error {
			if len(args) != 0 {
				t.Error("Expected no arguments, but got:", args)
			}
			return nil
		},
	}
	cli.AddCommand(*cmd)
	err := cli.Run([]string{"test"})
	if err != nil {
		t.Error("Expected command to run without error, but got:", err)
	}

}
