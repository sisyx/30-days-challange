package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
)

func main() {
	args := os.Args

	host := "scanme.nmap.org"
	if len(args) > 1 {
		host = args[1]
	}

	from := 1
	to := 30000

	if len(args) > 2 {
		from = string_to_int(args[2])
	}

	if len(args) > 3 {
		to = string_to_int(args[3])
	}

	var wg sync.WaitGroup
	for port := from; port <= to; port++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			address := fmt.Sprintf("%v:%d", host, port)
			conn, err := net.Dial("tcp", address)
			if err == nil {
				conn.Close()
			}
			fmt.Printf("Port %d open\n", port)
		}()
	}

	wg.Wait()
	fmt.Println("DONE!")
}

func string_to_int(text string) int {
	i, err := strconv.Atoi(text)
	if err != nil {
		// ... handle error
		panic(err)
	}
	return i
}
