package main

import (
	"bufio"
	"fmt"
	"os"
)

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		counts := countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			counts := countLines(f, counts)
			f.Close()

		}

	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	// Store every line in counts as:
	// key: "actual line contents", value: linecount
	for input.Scan() {
		line := input.Text()
		counts[file+" : "+line] = counts[file+" : "+line] + 1
	}
}
