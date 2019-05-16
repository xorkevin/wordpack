package linebreak

import (
	"fmt"
	"strings"
)

// splitWords returns an input string split into words, and the lengths of each
// word
func splitWords(paragraph string) ([]string, []int) {
	words := strings.Fields(paragraph)
	k := make([]int, 0, len(words))
	for _, i := range words {
		k = append(k, len(i))
	}
	return words, k
}

type (
	breakpoint struct {
		currcost  int
		prevcost  int
		prevbreak int
	}
)

// findBreakpointsKnuth takes in a list of word lengths and maximum width, and
// finds the optimal breakpoints minimizing the squares of line end spaces.
func findBreakpointsKnuth(wordLengths []int, width int) []int {
	return []int{}
}

// WrapParagraphs wraps stdin or an input file on a per paragraph basis. A
// paragraph is separated by two new lines.
func WrapParagraphs(args []string) {
	if len(args) > 0 {
		fmt.Println(args[0])
	} else {
		fmt.Println("stdin")
	}
}
