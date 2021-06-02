package main

import (
	"encoding/json"
	"github.com/kirigaikabuto/Golang1300Lessons/10/posts"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//send request and get response
	url := "http://jsonplaceholder.typicode.com/posts"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	//response body convert to bytes
	jsonData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	var postsData []posts.Post
	err = json.Unmarshal(jsonData, &postsData)
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, v := range postsData {
		v.ShowInfo()
	}
}
