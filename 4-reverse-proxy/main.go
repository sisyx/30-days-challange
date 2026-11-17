package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {

	args := os.Args
	proxy_from__port := 8000
	if len(args) >= 1 {
		proxy_from__port = string_to_int(args[1])
	}
	// proxy_from__host := "http://localhost"
	proxy_to__host := "http://localhost"
	if len(args) >= 2 {
		proxy_to__host = args[2]
	}
	proxy_to__port := 3000
	if len(args) >= 3 {
		proxy_to__port = string_to_int(args[3])
	}

	generalHandler := func(w http.ResponseWriter, req *http.Request) {
		slog.Info(time.Now().String(), req.Method, req.URL)
		// Send each request to port 3000
		newRequestUrl := fmt.Sprintf("%v:%v%v", proxy_to__host, proxy_to__port, req.URL)
		slog.Info("Sending request to %v", newRequestUrl)
		var resp *http.Response
		var err error
		switch req.Method {
		case "POST":
			resp, err = http.Post(newRequestUrl, req.Header.Get("Content-Type"), req.Body)
		default:
			resp, err = http.Get(newRequestUrl)
		}

		// set all headers from port 3000 to client response in port 8000
		for key, values := range resp.Header {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}

		// recive the response from :3000
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Internal Server Error"))
			return
		}

		// send the result to client on port 8000
		w.WriteHeader(resp.StatusCode)
		w.Write(body)
	}

	// handle all routes using generalHandler
	http.HandleFunc("/", generalHandler)

	// Listen to port 8000
	listen_port_string := fmt.Sprintf(":%v", proxy_from__port)
	http.ListenAndServe(listen_port_string, nil)

}

func string_to_int(text string) int {
	i, err := strconv.Atoi(text)
	if err != nil {
		// ... handle error
		panic(err)
	}
	return i
}
