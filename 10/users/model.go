package users

import "fmt"

type AddressModel struct {
	District   string
	Street     string
	HomeNumber string
}

func (a *AddressModel) getInfo() string {
	message := fmt.Sprintf("District:%s,Street:%s,HomeNumber:%s", a.District, a.Street, a.HomeNumber)
	return message
}

type User struct {
	Id       string
	Username string
	Password string
	Marks    []int
	Address  AddressModel
}

func (u *User) ShowInfo() {
	fmt.Println(u.Id, u.Username, u.Password, u.getAverageMark(), u.Address.getInfo())
}

func (u *User) getAverageMark() int {
	sumi := 0
	n := len(u.Marks)
	for _, v := range u.Marks {
		sumi += v
	}
	return sumi / n
}
