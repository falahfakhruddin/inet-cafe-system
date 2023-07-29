package command

import (
	"fmt"
	"strconv"

	"internet-cafe/pkg/model"
)

type CreateInternetCafe struct {
	pcNumber uint64
}

func (c *CreateInternetCafe) Run() error {
	_, err := model.NewInternetCafe(c.pcNumber)
	if err != nil {
		fmt.Println("error creating internet cafe")
		return err
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
