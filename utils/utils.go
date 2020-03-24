package utils

import (
	"fmt"
	"strings"
)

func LeftPadWithZeros(input string, size int) string {
	if len(input) <= 0 {
		return strings.Repeat("0", size)
	}

	if len(input) > size {
		return input
	}

	return fmt.Sprintf("%s%s", strings.Repeat("0", size - len(input)), input)

}
