package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

var filesModMap = map[string]time.Time{}

func main() {
	files := []string{}

	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if len(os.Args) > 1 {
		workDir = os.Args[1]
	}

	filepath.WalkDir(workDir, func(tmppath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		files = append(files, tmppath)
		return nil
	})

	for {
		time.Sleep(1000 * time.Millisecond)
		for _, file := range files {
			checkFile(file)
		}
	}
}

func checkFile(myfilepath string) {
	oldModTime := getModTime(myfilepath, true)
	modTime := getModTime(myfilepath, false)

	if oldModTime != modTime {
		fmt.Printf("File %v Changed at %v\n", myfilepath, modTime)
	}
	filesModMap[myfilepath] = getModTime(myfilepath, false)
}

func getModTime(tmpFilepath string, fromMap bool) time.Time {
	if fromMap {
		value, ok := filesModMap[tmpFilepath]
		if ok {
			return value
		}
	}
	info, err := os.Stat(tmpFilepath)
	if err != nil {
		panic(err)
	}

	filesModMap[tmpFilepath] = info.ModTime()

	return info.ModTime()
}
