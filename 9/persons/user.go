package persons

import "fmt"

type User struct {
	Name        string
	Username    string
	Password    string
	Age         int
	YearOfBirth int
}

func (u User) ShowName() {
	fmt.Println(u.Name)
}

func (u User) ShowUsernameAndPassword() {
	fmt.Println(u.Username, u.Password)
}

func (u *User) SetYearOfBirth() {
	u.YearOfBirth = 2021 - u.Age
}

func (u *User) SetName(name string) {
	u.Name = name
}
