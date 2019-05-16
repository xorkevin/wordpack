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
		cost int
		prev int
	}
)

func squareCost(x, w int) int {
	k := w - x
	return k * k
}

func lineCost(x, w int, last bool) int {
	if last {
		return 0
	}
	return squareCost(x, w)
}

// findBreakpointsKnuth takes in a list of word lengths and maximum width, and
// finds the optimal breakpoints minimizing the squares of line end spaces.
func findBreakpointsKnuth(wordLengths []int, width int) (int, []int) {
	last := len(wordLengths)
	dpCache := make([]breakpoint, 0, len(wordLengths)+1)
	dpCache = append(dpCache, breakpoint{
		cost: 0,
		prev: 0,
	})
	for i := range wordLengths {
		accWidth := wordLengths[i]
		minCost := lineCost(accWidth, width, i == last-1) + dpCache[i].cost
		minPrev := i
		for j := i - 1; j >= 0; j-- {
			accWidth += wordLengths[j] + 1
			if accWidth > width {
				break
			}
			c := lineCost(accWidth, width, i == last-1) + dpCache[j].cost
			if c < minCost {
				minCost = c
				minPrev = j
			}
		}
		dpCache = append(dpCache, breakpoint{
			cost: minCost,
			prev: minPrev,
		})
	}

	finalcost := dpCache[last].cost
	breakstack := []int{}
	for i := dpCache[last].prev; i > 0; i = dpCache[i].prev {
		breakstack = append(breakstack, i)
	}
	for i := 0; i < len(breakstack)/2; i++ {
		j := len(breakstack) - i - 1
		breakstack[i], breakstack[j] = breakstack[j], breakstack[i]
	}
	return finalcost, breakstack
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
