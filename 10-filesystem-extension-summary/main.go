package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

func main() {
	pathToAnalyze := "/mnt/30-day-challenge/"
	extCountMap := map[string]int{}
	filepath.WalkDir(pathToAnalyze, func(tmppath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			ext, hasExt := getExt(d.Name())
			if hasExt {
				extCountMap[ext] = extCountMap[ext] + 1
			}
		}
		return nil
	})
	for k, v := range extCountMap {
		fmt.Printf(".%v: %v\n", k, v)
	}
}

func getExt(filename string) (string, bool) {
	splitedFilename := strings.Split(filename, ".")

	if strings.Count(filename, ".") == 0 {
		return "", false
	}

	return splitedFilename[len(splitedFilename)-1], true
}
