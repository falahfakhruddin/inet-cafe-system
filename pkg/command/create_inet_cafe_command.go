package command

import (
	"fmt"
	"internet-cafe/pkg/dao"
	"strconv"

	"internet-cafe/pkg/model"
)

type CreateInternetCafe struct {
	InetCafeDao *dao.InternetCafeDao
	pcNumber    uint64
}

func (c *CreateInternetCafe) Run() error {
	inetCafe, err := model.NewInternetCafe(c.pcNumber)
	if err != nil {
		fmt.Println("error creating internet cafe")
		return err
	}

	err = c.InetCafeDao.Save(inetCafe)
	if err != nil {
		fmt.Println("error save internet cafe")
	}
	fmt.Println("success creating internet cafe")
	return nil
}

func (c *CreateInternetCafe) ParseArgs(args string) error {
	intVar, err := strconv.Atoi(args)
	if err != nil {
		return InvalidArgument
	}
	c.pcNumber = uint64(intVar)
	return nil
}
