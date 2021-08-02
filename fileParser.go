package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
)

func parseProject(s string, d fs.DirEntry, e error) error {
	if e != nil {
		return e
	}
	if !d.IsDir() {
		fmt.Println("Not a directory")
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
			fmt.Printf("TODO found on line: %d\n\t %s\n", line, scanner.Text())
		}
		line++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return "did it"
}
