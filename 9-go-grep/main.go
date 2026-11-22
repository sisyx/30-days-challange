package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/fatih/color"
)

func main() {

	phrase := getArgument(1, "Started")

	currentDir, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	fileToSearchIn := path.Join(currentDir, "test.log")

	foundLines, exists := findInFile(fileToSearchIn, phrase)

	if exists {
		for _, line := range foundLines {
			fmt.Println(strings.ReplaceAll(line, phrase, color.RedString(phrase)))
		}

	}

}

func findInFile(absoluteFilePath, phrase string) (lines []string, exists bool) {
	if phrase == "" {
		return nil, false
	}

	contentInByte, err := os.ReadFile(absoluteFilePath)
	if err != nil {
		panic(err)
	}
	for _, line := range strings.Split(string(contentInByte), "\n") {
		if strings.Contains(line, phrase) {
			lines = append(lines, line)
		}
	}

	return lines, len(lines) > 0
}

func getArgument(index int, defaultValue string) string {
	if index < 1 {
		return ""
	}

	if len(os.Args) > index {
		return os.Args[index]
	}

	return defaultValue
}
