package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	logs := get_logs()
	filtered_logs := filter_logs(logs)

	var ips []string
	var servers []string
	var hosts []string
	var requests []string
	lines := strings.Split(string(filtered_logs), "\n")

	for _, line := range lines {
		ip, err := get_entry(line, "client")
		if err != nil {
			ips = append(ips, "unknown")
		} else {
			ips = append(ips, ip)
		}
		server, err := get_entry(line, "server")
		if err != nil {
			servers = append(servers, "unknown")
		} else {
			servers = append(servers, server)
		}
		host, err := get_entry(line, "host")
		if err != nil {
			hosts = append(hosts, "unknown")
		} else {
			hosts = append(hosts, host)
		}
		request, err := get_entry(line, "request")
		if err != nil {
			requests = append(requests, "unknown")
		} else {
			requests = append(requests, request)
		}
	}

	fmt.Println("TOP ONES: ")
	fmt.Print("IP:\t")
	fmt.Println(find_top_one(ips))
	fmt.Print("SERVER:\t")
	fmt.Println(find_top_one(servers))
	fmt.Print("HOST:\t")
	fmt.Println(find_top_one(hosts))
	fmt.Print("REQUEST:\t")
	fmt.Println(find_top_one(requests))
}

func get_logs() []byte {
	currentDir, err := os.Getwd()
	check(err)

	logs_filename := "nginx.log"
	path := filepath.Join(currentDir, logs_filename)
	file, err := os.ReadFile(path)
	check(err)

	return file
}

func filter_logs(logs []byte) []byte {
	logs_string := string(logs)
	var filtered_lines []byte
	// we only need some specific lines
	for _, line := range strings.Split(logs_string, "\n") {
		contains_host := strings.Contains(line, "host")
		contains_request := strings.Contains(line, "request")
		contains_server := strings.Contains(line, "server")
		contains_client := strings.Contains(line, "client")
		contains_upstream := strings.Contains(line, "upstream")

		if contains_host && contains_request && contains_server && contains_client && contains_upstream {
			filtered_lines = append(filtered_lines, []byte(line)...)
			filtered_lines = append(filtered_lines, '\n')
		}

	}
	return filtered_lines
}

func get_entry(line, entery string) (string, error) {
	for _, phrase := range strings.Split(line, ",") {
		if strings.Contains(phrase, entery) && strings.Contains(phrase, ":") {
			return strings.Split(phrase, ":")[1], nil
		}
	}

	return "", errors.New("Entery not found")
}

func find_top_one(lst []string) string {
	count_map := map[string]int{}
	for _, item := range lst {
		count_map[item] = 0
	}

	for _, item := range lst {
		count_map[item] += 1
	}

	top_key := ""
	top_value := 0

	for k, v := range count_map {
		if top_value < v {
			top_key = k
			top_value = v
		}
	}

	return fmt.Sprintf("%v appears %v times", top_key, top_value)
}
