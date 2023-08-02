package command

import (
	"fmt"
	"internet-cafe/pkg/dao"
)

var CommandMapping = map[string]func(dependencies Dependencies) Command{
	"create_internet_cafe": NewCreateInternetCafeCommand,
	"status":               NewStatusInternetCafeCommand,
	"leave":                NewLeavePCCommand,
	"book":                 NewAllocatePCCommand,
	"get_pc_by_age":        NewGetPCByAgeCommand,
}

type Dependencies struct {
	InetCafeDao *dao.InternetCafeDao
}

func GetCommand(cmd string, dependencies Dependencies) (Command, error) {
	commandFunc, exists := CommandMapping[cmd]
	if !exists {
		return nil, fmt.Errorf("unknown command")
	}

	return commandFunc(dependencies), nil

}

func NewCreateInternetCafeCommand(dependencies Dependencies) Command {
	return &CreateInternetCafe{
		InetCafeDao: dependencies.InetCafeDao,
	}
}

func NewStatusInternetCafeCommand(dependencies Dependencies) Command {
	return &StatusInternetCafe{
		InetCafeDao: dependencies.InetCafeDao,
	}
}

func NewLeavePCCommand(dependencies Dependencies) Command {
	return &LeavePC{
		InetCafeDao: dependencies.InetCafeDao,
	}
}

func NewAllocatePCCommand(dependencies Dependencies) Command {
	return &AllocatePC{
		InetCafeDao: dependencies.InetCafeDao,
	}
}

func NewGetPCByAgeCommand(dependencies Dependencies) Command {
	return &GetPCByAge{
		InetCafeDao: dependencies.InetCafeDao,
	}
}
