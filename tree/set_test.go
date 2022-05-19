package tree_test

import (
	"testing"

	"github.com/weisrc/recordb/tree"
)

func TestSetMerge(t *testing.T) {
	root := tree.Root[int]()
	root.Set("a.example.com.", 0, true)
	root.Set("b.example.com.", 0, true)
	root.Set("c.example.com.", 0, true)
	Assert(t, root, `
	.example.com.
    |a!*
    |b!*
    |c!*
	`)
}

func TestSetFork(t *testing.T) {
	root := tree.Root[int]()
	root.Set("example.com.", 0, true)
	root.Set("com.", 0, true)
	Assert(t, root, `
	com.!*
    |example.!*
	`)
}

func TestSetExact(t *testing.T) {
	root := tree.Root[int]()
	root.Set("example.com.", 0, true)
	root.Set("example.com.", 0, true)
	Assert(t, root, `
	example.com.!*
	`)
}

func TestSetExtend(t *testing.T) {
	root := tree.Root[int]()
	root.Set("example.com.", 0, true)
	root.Set("a.example.com.", 0, true)
	root.Set("aa.example.com.", 0, true)
	Assert(t, root, `
	example.com.!*
	|a.!*
	||a!*
	`)
}
