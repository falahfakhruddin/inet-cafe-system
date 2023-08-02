package command

import (
	"fmt"
	"internet-cafe/pkg/dao"
	"strconv"
)

type LeavePC struct {
	InetCafeDao *dao.InternetCafeDao
	pcNumber    uint64
}

func (l *LeavePC) Run() error {
	inetCafe, err := l.InetCafeDao.Get()
	if err != nil {
		return err
	}

	success := inetCafe.LeavePC(l.pcNumber)
	if success {
		fmt.Println(fmt.Sprintf("pc number %d is free", l.pcNumber))
	}
	return nil
}

func (l *LeavePC) ParseArgs(args string) error {
	intVar, err := strconv.Atoi(args)
	if err != nil {
		return InvalidArgument
	}
	l.pcNumber = uint64(intVar)
	return nil
}
