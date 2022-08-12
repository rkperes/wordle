package dict

// I don't like this being here.
// It should be relevant for both dict and game, though.
// As dict is consumed by game, I'll leave it here for now.
const WordLen = 5

type Dict interface {
	Exists(word string) bool
	RandomWord() string
}
