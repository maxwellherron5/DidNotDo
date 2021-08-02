package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
)

func parseProject(path string, dir fs.DirEntry, e error) error {
	if e != nil {
		return e
	}
	if !dir.IsDir() {
		parseFile(path)
	}
	return nil
}

func parseFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	line := 1
	for scanner.Scan() {
		if strings.Contains(strings.ToLower(scanner.Text()), "todo") {
			fmt.Printf("\nTODO found on line %d with message: %s\n\n", line, scanner.Text())
		}
		line++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return "did it"
}
