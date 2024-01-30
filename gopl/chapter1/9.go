package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, copyStatusErr := io.Copy(os.Stdout, strings.NewReader("Response Status: \t"+resp.Status+"\n\n"))
		_, copyBodyErr := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if copyStatusErr != nil || copyBodyErr != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
