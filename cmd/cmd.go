package cmd

import "fatty/cmd/fatty"

type Command interface {
	Execute() error
}

type CommandType string
const (
	Fatty CommandType = "fatty"
)

var (
	commandLookup = map[CommandType]Command{
		Fatty: fatty.FattyCommand{},
	}
)

func RunCommand(str string) error {
	if handler, ok := commandLookup[CommandType(str)]; ok {
		return handler.Execute()
	}

	return nil
}