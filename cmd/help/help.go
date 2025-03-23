package help

import "fmt"

type HelpCommand struct{}

func (h HelpCommand) Execute() error {
	fmt.Printf("usage: 'fatty <command> [args]'\n")
	fmt.Printf("commands:\n")
	fmt.Printf("  fatty - this will create new accounts and attempt to generate a code\n")
	fmt.Printf("  accounts - mass generate accounts\n")
	fmt.Printf("  codes - mass generate codes\n")

	return nil
}