package linebreak

import (
	"bufio"
	"bytes"
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

func wrapParagraph(paragraph string, width int, w *bufio.Writer) error {
	words, wordLengths := splitWords(paragraph)
	_, breakstack := findBreakpointsKnuth(wordLengths, width)
	prev := 0
	for _, i := range breakstack {
		if _, err := w.WriteString(strings.Join(words[prev:i], " ")); err != nil {
			return err
		}
		if _, err := w.WriteString("\n"); err != nil {
			return err
		}
		prev = i
	}
	if _, err := w.WriteString(strings.Join(words[prev:], " ")); err != nil {
		return err
	}
	if _, err := w.WriteString("\n"); err != nil {
		return err
	}
	return nil
}

func splitParagraphs(data []byte, atEOF bool) (int, []byte, error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := bytes.Index(data, []byte("\n\n")); i >= 0 {
		if i == 0 {
			return 2, nil, nil
		}
		return i + 2, data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}
	return 0, nil, nil
}

// WrapParagraphs wraps stdin or an input file on a per paragraph basis. A
// paragraph is separated by two new lines.
func WrapParagraphs(width int, args []string) {
	if len(args) > 0 {
		fmt.Println(args[0])
	} else {
		fmt.Println("stdin")
	}
	fmt.Println(width)
}
