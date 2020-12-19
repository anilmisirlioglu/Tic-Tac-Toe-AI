package game

import (
	"fmt"
	"strings"
)

func Write(context string, nextLine int, prefix bool, multiplier int) {
	text := ""
	if nextLine == NextLineBefore || nextLine == NextLineMultiply {
		text += "\n"
	}

	if prefix {
		text += Prefix + " "
	}

	text += context

	if nextLine == NextLineBehind || nextLine == NextLineMultiply {
		text += strings.Repeat("\n", multiplier)
	}

	fmt.Printf(text)
}
