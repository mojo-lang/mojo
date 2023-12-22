package main

import (
	"os"
	"strings"
)

// MCC_TARGET_DIR

func main() {
	args := os.Args
	println(strings.Join(args, " "))
}
