package datastructures

import "errors"

func IsBalancedParen(s string) (bool, error) {
	// just count them in a variable?
	// first make sure there are no other ones?
	runes := []rune(s)

	for _, r := range runes {
		if !(r == 41 || r == 40) {
			return errors.New("this function only balances parens")
		}
	}

	return true, nil
}
