package symtable

import (
	"algorithms/types"
)

type Pair struct {
	Key types.Compare
	Value interface{}
}

type Node struct {
	Pair
	Next *Node
}

// NOTE: only support compare to another Pair
func (pair Pair)Less(other interface{}) bool{
	if o, ok := other.(Pair); ok {
		return pair.Key.Less(o.Key)
	} else {
		panic("map node only compare to another map pair")
	}
}

// NOTE: only support compare to another Pair
func (pair Pair)Compare(other interface{}) int {
	if o, ok := other.(Pair); ok {
		return pair.Key.Compare(o.Key)
	} else {
		panic("map node only compare to another map pair")
	}
}

type SymTable interface {
	Put(types.Compare, interface{})
	Get(types.Compare) interface{}
	Delete(types.Compare)
	Contains(types.Compare) bool
	IsEmpty() bool
	Size() int
	Keys() []types.Compare
}
