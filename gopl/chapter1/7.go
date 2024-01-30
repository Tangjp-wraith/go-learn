package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, copyErr := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if copyErr != nil {
			fmt.Fprintf(os.Stderr, "io.Copy err: %v", copyErr)
			os.Exit(1)
		}
	}
}
