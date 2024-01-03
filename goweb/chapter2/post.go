package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://www.shirdon.com/comment/add/"
	body := "{\"userId\":1,\"articleId\":1,\"comment\":\"这是一条评论\"}"
	response, err := http.Post(url, "application/x-www-form-urlencoded",
		bytes.NewBuffer([]byte(body)))
	if err != nil {
		fmt.Printf("err", err)
	}
	b, err := io.ReadAll(response.Body)
	fmt.Printf(string(b))
}
