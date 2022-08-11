package dict

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLocalExists(t *testing.T) {
	rq := require.New(t)
	d, err := NewLocalDict()
	rq.NoError(err)

	validWords := []string{
		"cigar",
		"shake",
		"navel",
	}
	for _, word := range validWords {
		rq.Truef(
			d.Exists(word),
			"expected word %q to exist",
			word,
		)
	}

	invalidWords := []string{
		"",
		"four",
		"letter",
		"cigra",
		"asdfg",
	}
	for _, word := range invalidWords {
		rq.Falsef(
			d.Exists(word),
			"expected word %q to not exist",
			word,
		)
	}
}

func TestLocalRandom(t *testing.T) {
	rq := require.New(t)
	d, err := NewLocalDict()
	rq.NoError(err)

	for i := 0; i < 1000; i++ {
		word := d.RandomWord()
		rq.Lenf(word, 5, "expected word %q to have length 5", word)
		rq.Truef(d.Exists(word), "expected word %q to exist", word)
	}
}
