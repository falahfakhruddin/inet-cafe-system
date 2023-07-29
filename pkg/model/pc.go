package model

type PC struct {
	ID   uint64
	User *User
}

func (p *PC) IsFree() bool {
	return p.User == nil
}

func (p *PC) FreeSlot() {
	p.User = nil
}
