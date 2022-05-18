package tree

import (
	"fmt"
	"strings"
)

type Tree struct {
	segment  string
	wild     bool
	Value    string
	children [38]*Tree
	parent   *Tree
}

func New(path string, value string, wild bool) *Tree {
	t := Empty(path)
	t.Value = value
	t.wild = wild
	return t
}

func Empty(path string) *Tree {
	t := new(Tree)
	t.segment = path
	return t
}

func (t *Tree) Path() string {
	var sb strings.Builder
	for {
		sb.WriteString(t.segment)
		if t = t.parent; t == nil {
			break
		}
	}
	return sb.String()
}

func (t *Tree) Remove(path string) *Tree {
	if node := t.Get(path); node != nil {
		node.parent.children[node.id()] = nil
		return node
	}
	return nil
}

func (t *Tree) PrettyPrint(tab int) {
	fmt.Printf("%s- %s\n", strings.Repeat("  ", tab), t.segment)
	for _, v := range t.children {
		if v != nil {
			v.PrettyPrint(tab + 1)
		}
	}
}
