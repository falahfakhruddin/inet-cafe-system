package command

import (
	"fmt"
	"strconv"
	"strings"

	"internet-cafe/pkg/model"
)

type AllocatePC struct {
	user *model.User
}

func (a *AllocatePC) Run() error {
	inetCafe, err := model.GetInternetCafe()
	if err != nil {
		fmt.Println("error allocate PC")
		return err
	}

	pc := inetCafe.AllocatePC(a.user)
	if pc == nil {
		fmt.Println("internet cafe is full book")
		return nil
	}
	fmt.Println(fmt.Sprintf("allocate PC number %d", pc.ID))
	return nil
}

func (a *AllocatePC) ParseArgs(args string) error {
	splitArgs := strings.Split(args, " ")
	if len(splitArgs) != 2 {
		return InvalidArgument
	}

	intAge, err := strconv.Atoi(splitArgs[1])
	if err != nil {
		return err
	}

	a.user = &model.User{
		Name: splitArgs[0],
		Age:  int64(intAge),
	}
	return nil
}
