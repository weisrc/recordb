package tree_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/weisrc/recordb/tree"
)

const N = 1000000
const I = 1000000

func BenchmarkTree(b *testing.B) {

	v, _ := uuid.NewRandom()
	s := tree.Hash(v.String())
	root := tree.New(s, "asdf", false)

	for i := 0; i < N; i++ {
		v, _ := uuid.NewRandom()
		root.Set(tree.Hash(v.String()), "asdf", false)
	}

	b.ResetTimer()

	for i := 0; i < I; i++ {
		if node := root.At(s, true); node == nil {
			b.Fail()
		}
	}
}

func BenchmarkMap(b *testing.B) {
	v, _ := uuid.NewRandom()
	s := v.String()
	root := make(map[string]string)
	root[s] = "Hello"

	for i := 0; i < N; i++ {
		v, _ := uuid.NewRandom()
		root[v.String()] = "Hello"
	}

	b.ResetTimer()

	for i := 0; i < I; i++ {
		if _, ok := root[s]; !ok {
			b.Fail()
		}
	}
}
