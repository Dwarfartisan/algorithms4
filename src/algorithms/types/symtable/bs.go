package symtable

import (
	"algorithms/types"
)


type BsMap struct{
	buffer []*Pair
	size int
}

func NewBsMap(capacity int) (bst *BsMap){
	var buffer = make([]*Pair, capacity)
	var ret = &BsMap{buffer, 0}
	return ret
}

func (bst *BsMap)Put(key types.Compare, value interface{}){
	var r,ok = bst.rankKey(key)
	if ok {
		bst.buffer[r].Value = value
	} else {
		bst.insert(r, &Pair{key, value})
	}
}

func (bst BsMap)rank(node Node) int {
	if bst.size == 0 {
		return 0
	}
	if bst.Max().Compare(node) < 0 {
		return bst.Size()
	}
	var (
		lo = 0 
		hi = bst.Size()
	)

	for lo<=hi {
		var mid = lo + (hi-lo)/2
		var cmp = node.Compare(bst.buffer[mid])
		if cmp < 0 {
			hi = mid-1
		} else if cmp > 0 {
			lo = mid+1
		} else {
			return mid
		}
	}
	return lo
}

func (bst BsMap)rankKey(key types.Compare) (int, bool) {
	if bst.size == 0 {
		return 0, false
	}
	if bst.Min().Key.Compare(key) > 0 {
		return 0, false
	}
	if bst.Max().Key.Compare(key) < 0 {
		return bst.Size(), false
	}
	var (
		lo = 0 
		hi = bst.Size()
	)

	for lo<=hi {
		var mid = lo + (hi-lo)/2
		var cmp = key.Compare(bst.buffer[mid].Key)
		if cmp < 0 {
			hi = mid-1
		} else if cmp > 0 {
			lo = mid+1
		} else {
			return mid, true
		}
	}
	return lo, false
}

func (bst BsMap)Get(key types.Compare)(interface{}) {
	if bst.Size() == 0 {
		return nil
	}
	var r,ok = bst.rankKey(key)
	if !ok {
		return nil
	} else {
		var node = bst.buffer[r]
		return node.Value
	}
}

func (bst *BsMap)Delete(key types.Compare) {
	var r, ok = bst.rankKey(key)
	if ok {
		copy(bst.buffer[r:], bst.buffer[r+1:])
		bst.size -= 1
	}
}

func (bst BsMap)Contains(key types.Compare) bool {
	var _, ok = bst.rankKey(key)
	return ok
}

func (bst BsMap) Size() int{
	return bst.size
}

func (bst BsMap) Capacity() int {
	return len(bst.buffer)
}

func (bst BsMap)Min() *Pair {
	if bst.size == 0 {
		return nil
	}
	return bst.buffer[0]
}

func (bst BsMap)Max() *Pair {
	if bst.size == 0 {
		return nil
	}
	return bst.buffer[bst.Size()-1]
}

func (bst BsMap)Select(k int) *Pair {
	return bst.buffer[k]
}

func (bst *BsMap)DeleteMin(){
	if bst.size == 0 {
		return
	}
	bst.buffer = bst.buffer[1:]
	bst.size -= 1
}

func (bst *BsMap)DeleteMax() {
	if bst.size == 0 {
		return
	}
	bst.buffer = bst.buffer[bst.size-1:]
	bst.size -=1
}

func (bst *BsMap)insert(index int, node *Pair){
	var buffer = bst.buffer
	if bst.Size() == bst.Capacity() {
		bst.buffer = make([]*Pair, bst.Capacity()*2)
		copy(bst.buffer, buffer[:index])
	} 		
	copy(bst.buffer[index+1:], buffer[index:])
    bst.buffer[index] = node
	bst.size+=1
}

func (bst BsMap)Keys() []types.Compare {
	var ret = make([]types.Compare, bst.size)
	for idx, node := range bst.buffer[:bst.size] {
		ret[idx] = node.Key
	}
	return ret
}

func (bst BsMap)IsEmpty()bool{
	return bst.size==0
}
