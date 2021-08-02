package main

import (
	"path/filepath"
)

func main() {
	filepath.WalkDir("/test", parseProject)
}
