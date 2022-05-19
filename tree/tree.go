package tree

import (
	"fmt"
	"strings"
)

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

func Empty[T any](key string) *Tree[T] {
	t := new(Tree[T])
	t.segment = key
	return t
}

func (t *Tree[T]) Path() string {
	var sb strings.Builder
	for {
		sb.WriteString(t.segment)
		if t = t.parent; t == nil {
			break
		}
	}
	return sb.String()
}

func (t *Tree[T]) PrettyPrint(tab int) {
	fmt.Printf("%s- %s\n", strings.Repeat("  ", tab), t.segment)
	for _, v := range t.children {
		if v != nil {
			v.PrettyPrint(tab + 1)
		}
	}
}
