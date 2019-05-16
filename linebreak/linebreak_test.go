package linebreak

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_splitWords(t *testing.T) {
	assert := assert.New(t)

	words, wordLengths := splitWords(`Lorem ipsum dolor sit amet,
	consectetur adipiscing elit.`)
	assert.Equal([]string{"Lorem", "ipsum", "dolor", "sit", "amet,", "consectetur", "adipiscing", "elit."}, words, "words should split on any space character")
	assert.Equal([]int{5, 5, 5, 3, 5, 11, 10, 5}, wordLengths, "word lengths should match their corresponding word lengths")
}

func Test_findBreakpointsKnuth(t *testing.T) {
	assert := assert.New(t)

	cost, breakstack := findBreakpointsKnuth([]int{3, 2, 2, 5}, 6)
	assert.Equal(10, cost, "cost should be sum of squares discounting the last line")
	assert.Equal([]int{1, 3}, breakstack, "breakstack should not be greedy")
}
