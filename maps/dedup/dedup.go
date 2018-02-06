package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// This program is a simple way to create a set
	seen := map[string]bool{}
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}
	if err := input.Err(); err != nil {
		log.Fatal(err)
	}
}
