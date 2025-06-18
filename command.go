package cli

import "flag"

// Define the Command structure and its methods
type Command struct {
	Name        string
	Description string
	Flags       []*flag.Flag
	Run         func(args []string) error
}
