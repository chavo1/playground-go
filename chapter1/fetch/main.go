// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const url_prefix = "http://" // we declare the "url_prefix"

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, url_prefix) { // 1.8 we adding the "url_prefix" in each url if missing
			url = url_prefix + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("HTTP status: %s\n", resp.Status) // 1.9 HTTP status: 200 OK - resp.StatusCode is an actual integer with the code (200)

		_, err = io.Copy(os.Stdout, resp.Body) // 1.7 we use io.Copy function call to reads from src and writes dst(os.Stdout). Also we are checking for error io.Copy.
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: copying: %v\n", err)
			os.Exit(1)
		}

	}
}
