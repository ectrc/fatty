package main

import (
	"fatty/cmd"
	"fmt"
	"os"
)

func main() {
	os.Args = append(os.Args, "fatty")

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <command> [args]\n", os.Args[0])
		return
	}

	if err := cmd.RunCommand(os.Args[1]); err != nil {
		fmt.Println(err)
	}
}