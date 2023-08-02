package model

import "fmt"

type InternetCafe struct {
	Capacity uint64
	ListPC   []*PC
}

var (
	savedInternetCafe *InternetCafe
)

func NewInternetCafe(capacity uint64) (*InternetCafe, error) {
	if savedInternetCafe != nil {
		return nil, fmt.Errorf("internet cafe already exist")
	}

	var index uint64 = 1
	inetCafe := new(InternetCafe)
	inetCafe.Capacity = capacity
	inetCafe.ListPC = make([]*PC, 0, capacity)
	for index <= capacity {
		pc := &PC{
			ID: index,
		}
		inetCafe.ListPC = append(inetCafe.ListPC, pc)
		index++
	}

	return inetCafe, nil
}

func (i *InternetCafe) AllocatePC(user *User) *PC {
	for _, pc := range i.ListPC {
		if pc.IsFree() {
			pc.User = &User{
				Name: user.Name,
				Age:  user.Age,
			}
			return pc
		}
	}
	return nil
}

func (i *InternetCafe) LeavePC(pcID uint64) bool {
	for _, pc := range i.ListPC {
		if pc.ID == pcID {
			pc.FreeSlot()
			return true
		}
	}
	return false
}

func (i *InternetCafe) GetPCIdByAge(age int64) []uint64 {
	var listPCId []uint64
	for _, pc := range i.ListPC {
		if pc.User != nil && pc.User.isSameAge(age) {
			listPCId = append(listPCId, pc.ID)
		}
	}

	return listPCId
}
