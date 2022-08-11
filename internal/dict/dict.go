package dict

type Dict interface {
	Exists(word string) bool
	RandomWord() string
}
