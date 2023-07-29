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
	inetCafe.Save()

	return inetCafe, nil
}

func GetInternetCafe() (*InternetCafe, error) {
	if savedInternetCafe == nil {
		return nil, fmt.Errorf("internet cafe not found")
	}
	return savedInternetCafe, nil
}

func (p *InternetCafe) Save() {
	savedInternetCafe = p
}

func (p *InternetCafe) AllocatePC(car *User) *PC {
	for _, pc := range p.ListPC {
		if pc.IsFree() {
			pc.User = &User{
				Name: car.Name,
				Age:  car.Age,
			}
			return pc
		}
	}
	return nil
}

func (p *InternetCafe) LeavePC(pcID uint64) bool {
	for _, pc := range p.ListPC {
		if pc.ID == pcID {
			pc.FreeSlot()
			return true
		}
	}
	return false
}
