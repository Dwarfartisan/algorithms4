package symtable

//ListMap base on list create a list map struct
import ("algorithms/types")

type ListMap struct {
	first *Node
}

func (l ListMap) Get(key types.Compare) interface{} {
	var node = l.first
	var ret interface{} = nil
	for {
		if node == nil {
			break
		}
		if node.Key.Compare(key)==0 {
			ret = node.Value
			break
		} else {
			node = node.Next
		}
	}
	return ret
}

func (l *ListMap) Put(key types.Compare, value interface{}){
	if l.first == nil {
		l.first = &Node{key, value, nil}
		return
	}
	var node = l.first
	for {
		if node.Key.Compare(key)==0 {
			node.Value = value
			return
		} else {
			if node.Next == nil {
				break
			} else {
				node = node.Next
			}
		}
	}
	node.Next = &Node{key, value, nil}
}

func (l *ListMap)Delete(key types.Compare) {
	var node = l.first
	var prev *Node = nil
	if node.Key.Compare(key) == 0 {
		l.first = node.Next
	}
	for {
		prev = node
		var node = node.Next
		if node.Key.Compare(key) == 0 {
			prev.Next = node.Next
			return
		}
		if node == nil {
			return
		}
	}
}

func (l ListMap)Contains(key types.Compare) bool {
	var node = l.first
	for {
		if node == nil {
			return false
		}
		if node.Key.Compare(key) == 0 {
			return true
		} else {
			node = node.Next
		}
	}
}

func (l ListMap) isEmpty() bool {
	return l.first == nil
}

func (l ListMap) Size() int {
	var count = 0
	var node = l.first
	for {
		if node == nil {
			return count
		}
		count++
		node = node.Next
	}
}

func (l ListMap)Keys() []types.Compare {
	var keys = make([]types.Compare, l.Size())
	var node = l.first
	for {
		if node == nil {
			return keys
		}
		keys = append(keys, node.Key)
		node = node.Next
	}
}


func NewListMap() *ListMap {
	return &ListMap{nil}
}

