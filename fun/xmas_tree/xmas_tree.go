package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	PrintTree(7)
}

func PrintTree(length int) {
	width := 3
	spaces := DrawSpaces(length, 0)
	// Topper
	fmt.Printf("%s @\n", spaces)
	for i := 1; i < (length*2); i++ {
		if i % 2 == 0 {
			width += 2
			spaces = DrawSpaces(length, i/2)
		}
		fmt.Printf("%s/%s\\\n", spaces, DrawAsterixes(width-2))
	}
	fmt.Printf("%s", DrawFooter(length))
}

func DrawSpaces(length int, i int) string {
	spaces := (length - 1) - i
	if spaces > 0 {
		return strings.Repeat(" ", (length - 1) - i)
	}
	return ""
}

func DrawFooter(length int) string {
	base := fmt.Sprintf("%s| |\n", DrawSpaces(length, 0))
	return strings.Repeat(base, 3)
}

func DrawAsterixes(val int) string {
	asterixes := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < val; i++ {
		if rand.Intn(20) == 1 {
			asterixes += "0"
		} else {
			asterixes += "*"
		}
	}
	return asterixes
}
