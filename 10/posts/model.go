package posts

import "fmt"

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (p *Post) ShowInfo() {
	message := fmt.Sprintf("Id:%d,Title:%s", p.Id, p.Title)
	fmt.Println(message)
}
