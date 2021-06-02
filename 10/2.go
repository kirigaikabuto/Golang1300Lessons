package main

import (
	"encoding/json"
	"fmt"
	"github.com/kirigaikabuto/Golang1300Lessons/10/posts"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//send request and get response
	url := "http://jsonplaceholder.typicode.com/posts/1"
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
	//bytes from response convert to needed struct
	var p1 posts.Post
	err = json.Unmarshal(jsonData, &p1)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(p1.Title)
	fmt.Println(p1.Body)
}
