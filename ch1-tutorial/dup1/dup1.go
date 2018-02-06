package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	dupLineCounts := countLines(os.Stdin)

	for line, n := range dupLineCounts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(r io.Reader) map[string]int {
	counts := make(map[string]int)
	input := bufio.NewScanner(r)

	for input.Scan() {
		counts[input.Text()]++
		// This is equivalent to:
		// line := input.Text()
		// counts[line] = counts[line] + 1
	}

	return counts
}
