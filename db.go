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
		return root.Get(name, wild)
	}
	return
}

func (db *DB) Remove(typ uint16, name string) {
	if root := db.trees[types.Hash(typ)]; root != nil {
		if node := root.Remove(name); node != nil {
			record := node.Value
			if record.Prev != nil {
				record.Prev.Next = record.Next
			} else {
				db.heads[name] = record.Next
			}
			if record.Next != nil {
				record.Next.Prev = record.Prev
			}
		}
	}
}

func (db *DB) Add(value Record) {
	hash := types.Hash(value.Type)
	root := db.trees[hash]
	name := value.Name
	ptr := &value
	if root == nil {
		root = tree.New(name, ptr, value.Wild)
		db.trees[hash] = root
	} else {
		root.Set(name, ptr, value.Wild)
	}
	old := db.heads[name]
	db.heads[name], ptr.Next = ptr, old
	if old != nil {
		old.Prev = ptr
	}
}

func (db *DB) List(name string) *Record {
	return db.heads[name]
}
