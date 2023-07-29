package command

import (
	"fmt"
	"strconv"

	"internet-cafe/pkg/model"
)

type GetPCByAge struct {
	age int64
}

func (g *GetPCByAge) Run() error {
	inetCafe, err := model.GetInternetCafe()
	if err != nil {
		return err
	}

	listID := inetCafe.GetPCIdByAge(g.age)
	if len(listID) <= 0 {
		fmt.Println("not pc found for age ", g.age)
		return nil
	}

	fmt.Println(listID)
	return nil
}

func (g *GetPCByAge) ParseArgs(args string) error {
	intVar, err := strconv.Atoi(args)
	if err != nil {
		return InvalidArgument
	}
	g.age = int64(intVar)
	return nil
}
