package command

import (
	"fmt"
	"strings"
)

type Command interface {
	Run() error
	ParseArgs(string) error
}

type MainCommand struct {
	Args    string
	Command Command
}

var (
	InvalidArgument = fmt.Errorf("invalid argument")
)

func InitCommand(commandCli string) (*MainCommand, error) {
	cmds := strings.SplitN(commandCli, " ", 2)
	command, exists := CommandMapping[cmds[0]]
	if !exists {
		return nil, fmt.Errorf("unknown command")
	}
	mainCommand := &MainCommand{
		Command: command,
	}

	if len(cmds) > 1 {
		mainCommand.Args = cmds[1]
	}

	return mainCommand, nil
}

func (m *MainCommand) Run() error {
	err := m.Command.ParseArgs(m.Args)
	if err != nil {
		return err
	}
	return m.Command.Run()
}
