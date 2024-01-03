package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	url := "https://www.shirdon.com/comment/update/"
	payload := strings.NewReader("{\"userId\":1,\"articleId\":1,\"comment\":\"这是一条评论\"}")
	req, _ := http.NewRequest("PUT", url, payload)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	fmt.Println(res)
	fmt.Println(string(body))
}
