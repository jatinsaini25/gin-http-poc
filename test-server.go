// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"strings"
// )

// func main() {

// 	url := "http://localhost:8000/posts"
// 	method := "GET"

// 	payload := strings.NewReader("")

// 	client := &http.Client{}
// 	req, err := http.NewRequest(method, url, payload)

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	req.Header.Add("Authorization", "Basic cHJhZ2ltdGVjaDpyZXZpZXdz")

// 	res, err := client.Do(req)
// 	defer res.Body.Close()
// 	body, err := ioutil.ReadAll(res.Body)

// 	fmt.Println(string(body))
// }
