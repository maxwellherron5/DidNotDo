package main

import (
	"path/filepath"
)

// todo : add docstrings
func main() {
	filepath.WalkDir("../DidNotDo/", parseProject)
}
