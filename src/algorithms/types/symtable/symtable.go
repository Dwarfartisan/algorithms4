package symtable

import (
	"algorithms/types"
)

type Node struct {
	Key types.Compare
	Value interface{}
	Next *Node
}

type SymTable interface {
	Put(types.Compare, interface{})
	Get(types.Compare) interface{}
	Delete(types.Compare)
	Contains(types.Compare) bool
	isEmpty() bool
	Size() int
	Keys() []types.Compare
}
