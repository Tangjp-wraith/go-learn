package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://baidu.com")
	if err != nil {
		fmt.Printf("err", err)
	}
	closer := resp.Body
	bytes, err := io.ReadAll(closer)
	if err != nil {
		fmt.Printf("err", err)
	}
	fmt.Println(string(bytes))
}
