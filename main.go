package main

import (
	"fatty/cmd"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: 'fatty <command> [args]'\nnote: use 'fatty help' to get a list of all of the possible commands\n")
		return
	}

	if err := cmd.RunCommand(os.Args[1]); err != nil {
		fmt.Println(err)
	}
}