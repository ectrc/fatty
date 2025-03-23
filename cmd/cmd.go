package cmd

import (
	"fatty/cmd/accounts"
	"fatty/cmd/codes"
	"fatty/cmd/fatty"
	"fatty/cmd/help"
	"fatty/services/config"
	"fmt"
	"time"
)

type Command interface {
	Execute() error
}

type CommandType string
const (
	Help CommandType = "help"
	Fatty CommandType = "fatty"
	Accounts CommandType = "accounts"
	Codes CommandType = "codes"
)

var (
	commandLookup = map[CommandType]Command{
		Help: help.HelpCommand{},
		Fatty: fatty.FattyCommand{},
		Accounts: accounts.AccountsGeneratorCommand{},
		Codes: codes.CodeGeneratorCommand{},
	}
)

func RunCommand(str string) error {
	config := config.Config()

	if config.ENABLE_START_TIME {
		fmt.Printf("Waiting until '%s' to start.\n", config.START_TIME)
		time.Sleep(time.Until(config.START_TIME))
	}

	if handler, ok := commandLookup[CommandType(str)]; ok {
		return handler.Execute()
	}

	return fmt.Errorf("unknown command: %s", str)
}