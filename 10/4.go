package main

import (
	"bytes"
	"encoding/json"
	"github.com/kirigaikabuto/Golang1300Lessons/10/posts"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "http://jsonplaceholder.typicode.com/posts"
	p1 := posts.Post{
		UserId: 12,
		Title:  "hello from yerassyl",
		Body:   "1321212312",
	}
	jsonData, err := json.Marshal(p1)
	if err != nil {
		log.Fatal(err)
		return
	}
	body := bytes.NewReader(jsonData)
	response, err := http.Post(url, "application/json", body)
	if err != nil {
		log.Fatal(err)
		return
	}
	jsonResponseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	p2 := posts.Post{}
	err = json.Unmarshal(jsonResponseData, &p2)
	if err != nil {
		log.Fatal(err)
		return
	}
	p2.ShowInfo()
}
