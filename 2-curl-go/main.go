package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Conf struct {
	maxRetires uint
}

func main() {
	conf := Conf{
		maxRetires: 3,
	}

	method := "Get"
	if len(os.Args) >= 2 {
		method = os.Args[1]
	}

	url := "https://jsonplaceholder.typicode.com/posts/1"
	if len(os.Args) >= 3 {
		url = os.Args[2]
	}
	for run := range int(conf.maxRetires) {
		fmt.Println()
		fmt.Printf("RUN %d\n", run+1)
		body, err := execute(method, url)

		if err != nil {
			fmt.Println(err.Error())
			if run >= int(conf.maxRetires)-1 {
				fmt.Printf("Failed All %v attempts\n", conf.maxRetires)
				break
			} else {
				fmt.Println("Retring...")
			}
		} else {
			fmt.Println(body)
			break
		}
	}
}

func execute(method, url string) (string, error) {
	var resp *http.Response
	var err error

	switch strings.ToLower(method) {
	case "post":
		resp, err = http.Post(url, "application/json", nil)
	default:
		resp, err = http.Get(url)
	}

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	return string(body), nil
}
