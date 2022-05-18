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
	s := v.String()
	root := tree.Empty(s)
	root.Value = "Hello"

	for i := 0; i < N; i++ {
		v, _ := uuid.NewRandom()
		root.Set(v.String(), "Hello", false)
	}

	b.ResetTimer()

	for i := 0; i < I; i++ {
		if node := root.Get(s); node == nil {
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
