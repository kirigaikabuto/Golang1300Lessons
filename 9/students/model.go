package students

import "fmt"

type Student struct {
	Username    string
	Password    string
	FirstName   string
	LastName    string
	AverageMark int
	Marks       []int
}

func (st Student) ShowInfo() {
	fullName := st.getFullName()
	avg := st.getAverageMark()
	info := fmt.Sprintf("Username:%s;FullName:%s;AverageMark:%d", st.Username, fullName, avg)
	fmt.Println(info)
}

func (st Student) getFullName() string {
	return st.FirstName + " " + st.LastName
}

func (st Student) getAverageMark() int {
	sumi := 0
	for i := 0; i < len(st.Marks); i++ {
		sumi += st.Marks[i]
	}
	avg := sumi / len(st.Marks)
	return avg
}
