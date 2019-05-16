package linebreak

import (
	"bufio"
	"bytes"
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

func Test_wrapParagraph(t *testing.T) {
	assert := assert.New(t)

	b := bytes.Buffer{}
	w := bufio.NewWriter(&b)
	err := wrapParagraph("AAA BB CC DDDDD", 6, w)
	assert.Nil(err, "wrapping a paragraph should not error on write")
	err = w.Flush()
	assert.Nil(err, "writing to a buffer should not error on write")
	s := b.String()
	assert.Equal("AAA\nBB CC\nDDDDD\n", s, "string should not be greedily wrapped")
}

func Test_splitParagraphs(t *testing.T) {
	assert := assert.New(t)

	b := bytes.Buffer{}
	b.WriteString("A B\nC\nD E\n\nF G\n\n\n\nH I\nJ\n\nK\n")
	scanner := bufio.NewScanner(&b)
	scanner.Split(splitParagraphs)
	k := make([]string, 0, 4)
	for scanner.Scan() {
		k = append(k, scanner.Text())
	}
	assert.Equal("A B\nC\nD E", k[0], "newline should not create a new paragraph")
	assert.Equal("F G", k[1], "paragraphs can be a single line")
	assert.Equal("H I\nJ", k[2], "more than 2 newlines do not create a new paragraph")
	assert.Equal("K\n", k[3], "end of file is not discarded")
}
