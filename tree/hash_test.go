package tree_test

import (
	"testing"

	"github.com/weisrc/recordb/tree"
)

func TestHashCaseInsentivity(t *testing.T) {
	root := tree.Root[int]()
	root.Set(tree.Hash("a.example.com."), 0, true)
	root.Set(tree.Hash("b.example.com."), 0, true)
	root.Set(tree.Hash("c.example.com."), 0, true)
	root.Set(tree.Hash("C.example.com."), 0, true)
	Assert(t, root, `
	.example.com.
    |a!*
    |b!*
    |c!*
	`)
}
