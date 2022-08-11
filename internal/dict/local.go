package dict

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

//go:embed local.json
var localBytes []byte

func NewLocalDict() (Dict, error) {
	var words []string
	if err := json.Unmarshal(localBytes, &words); err != nil {
		return nil, fmt.Errorf("while parsing local dictionary: %w", err)
	}

	if len(words) == 0 {
		return nil, errors.New("unexpected empty local dictionary")
	}

	indexedWords := make(map[string]struct{}, len(words))
	for _, word := range words {
		indexedWords[word] = struct{}{}
	}

	return &local{
		rd:      rand.New(rand.NewSource(time.Now().Unix())),
		words:   words,
		indexed: indexedWords,
	}, nil
}

type local struct {
	rd      *rand.Rand
	words   []string
	indexed map[string]struct{}
}

func (l local) Exists(word string) bool {
	_, ok := l.indexed[word]
	return ok
}

func (l local) RandomWord() string {
	idx := l.rd.Int() % len(l.words)
	return l.words[idx]
}
