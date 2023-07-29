package model

type User struct {
	Name string
	Age  int64
}

func (u *User) isSameAge(age int64) bool {
	return u.Age == age
}
