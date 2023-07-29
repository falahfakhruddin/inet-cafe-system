package model

type PC struct {
	ID   uint64
	User *User
}

func (s *PC) IsFree() bool {
	return s.User == nil
}

func (s *PC) FreeSlot() {
	s.User = nil
}
