package iteration

import "strings"

func Repeat(character string, number int) string {
	var result strings.Builder
	for i := 0; i < number; i++ {
		result.WriteString(character)
	}
	return result.String()
}
