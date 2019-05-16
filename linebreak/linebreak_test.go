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
