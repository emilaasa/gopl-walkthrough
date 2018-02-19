package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("unicoder takes exactly one argument, a valid unicode char")
		os.Exit(1)
	}

	fmt.Println(Unicoder(os.Args[1]))

}

// Unicoder takes a string returns the unicode code point in decimal (rune / int32)
func Unicoder(s string) rune {
	r, _ := utf8.DecodeRuneInString(s)
	return r
}
