package command

import (
	"fmt"

	"internet-cafe/pkg/model"
)

type StatusInternetCafe struct {
	arg string
}

func (s *StatusInternetCafe) Run() error {
	inetCafe, err := model.GetInternetCafe()
	if err != nil {
		return err
	}

	for _, pc := range inetCafe.ListPC {
		fmt.Println(fmt.Sprintf("PC: %d User: %+v", pc.ID, pc.User))
	}
	return nil
}

func (s *StatusInternetCafe) ParseArgs(args string) error {
	return nil
}
