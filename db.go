package recordb

import (
	"github.com/weisrc/recordb/tree"
	"github.com/weisrc/recordb/types"
)

type DB struct {
	trees [87]*tree.Tree[*Record]
	heads map[string]*Record
}

func New() *DB {
	s := new(DB)
	s.heads = make(map[string]*Record)
	return s
}

func (db *DB) Get(typ uint16, name string, wild bool) (value *Record, ok bool) {
	if root := db.trees[types.Hash(typ)]; root != nil {
		return root.Get(tree.Hash(name), wild)
	}
	return
}

func (db *DB) Remove(typ uint16, name string) {
	segment := tree.Hash(name)
	if root := db.trees[types.Hash(typ)]; root != nil {
		if node := root.Remove(segment); node != nil {
			record := node.Value
			if record.Prev != nil {
				record.Prev.Next = record.Next
			} else {
				db.heads[segment] = record.Next
			}
			if record.Next != nil {
				record.Next.Prev = record.Prev
			}
		}
	}
}

func (db *DB) Add(value Record) {
	typ := types.Hash(value.Type)
	root := db.trees[typ]
	segment := tree.Hash(value.Name)
	ptr := &value
	if root == nil {
		root = tree.New(segment, ptr, value.Wild)
		db.trees[typ] = root
	} else {
		root.Set(segment, ptr, value.Wild)
	}
	old := db.heads[segment]
	db.heads[segment], ptr.Next = ptr, old
	if old != nil {
		old.Prev = ptr
	}
}

func (db *DB) All(name string) *Record {
	return db.heads[tree.Hash(name)]
}
