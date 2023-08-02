package command

import (
	"fmt"
	"internet-cafe/pkg/dao"
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
	// Parse CLI Command
	cmds := strings.SplitN(commandCli, " ", 2)

	// Initialize Dependencies
	dep := InitializeDependencies()

	// Resolve CLI command to machine state command
	command, err := GetCommand(cmds[0], dep)
	if err != nil {
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

func InitializeDependencies() Dependencies {
	return Dependencies{
		InetCafeDao: dao.NewInternetCafeDao(),
	}
}
