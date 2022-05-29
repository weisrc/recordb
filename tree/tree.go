package tree

type Tree[T any] struct {
	Value    T
	Wild     bool
	leaf     bool
	segment  string
	count    byte
	children [38]*Tree[T]
	parent   *Tree[T]
}

func New[T any](key string, value T, wild bool) *Tree[T] {
	t := Empty[T](key)
	t.Value = value
	t.Wild = wild
	t.leaf = true
	return t
}

func Root[T any]() *Tree[T] {
	return new(Tree[T])
}

func Empty[T any](key string) *Tree[T] {
	t := Root[T]()
	t.segment = key
	return t
}
