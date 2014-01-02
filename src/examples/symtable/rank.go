package main

import (
	"fmt"
	"algorithms/types"
	"algorithms/types/symtable"
	"flag"
)

func main(){
	flag.Parse()
	var words = flag.Args()
	var st = symtable.NewListMap()
	FrequencyCounter(st, words)
}

func FrequencyCounter(st symtable.SymTable, words []string) {
	for _, w := range words {
		var word = types.String(w)
		if !st.Contains(word) {
			st.Put(word, 1)
		} else {
			var v = st.Get(word)
			value, ok := v.(int)
			if ok {
				st.Put(word, value+1)
			}
		}
	}
	var max_word = types.String("")
	var max_count = 0
	for _, w := range st.Keys() {
		var word, ok = w.(types.String)
		if ok {
			v := st.Get(word)
			value, ok := v.(int) 
			if ok {
				if max_count < value {
					max_word = word
					max_count = value
				}
			}
		}
	}
	fmt.Printf("max: %v value: %v\n", max_word, max_count)
}
