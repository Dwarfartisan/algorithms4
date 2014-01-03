package types

import "fmt"

type Less interface {
	Less(other interface{}) bool
}

type Compare interface {
	Less
	Compare(interface{}) int
}

type Numeric interface {
	Numeric() float64
}

type Integer int
func (i Integer)Numeric() float64{
	return float64(int(i))
}
func (i Integer)Less(other interface{}) bool{
	switch o := other.(type) {
	case int:
		return int(i)<o
	case float64:
		return i.Numeric()<o
	default:
		if d, ok :=o.(Numeric); ok{
			return i.Numeric()<d.Numeric()
		} else {
			var message = fmt.Sprintf("%v can't compare with %v", i, o)
			panic(message)
		}
	}
}
func(i Integer)Compare(other interface{}) int {
	switch o := other.(type) {
	case int:
		return CompareII(int(i), o)
	case float64:
		return CompareFF(i.Numeric(), o)
	default:
		if d, ok :=o.(Numeric); ok{
			return CompareFF(i.Numeric(), d.Numeric())
		} else {
			var message = fmt.Sprintf("%v can't compare with %v", i, o)
			panic(message)
		}
	}
}

func CompareII(x int, y int) int {
	if x>y {
		return 1
	}
	if x<y {
		return -1
	}
	return 0
}

type Float float64 
func (f Float)Numeric() float64{
	return float64(f)
}
func (f Float)Less(other interface{}) bool{
	switch o := other.(type) {
	case int:
		return f.Numeric()<float64(o)
	case float64:
		return f.Numeric()<o
	default:
		if d, ok :=o.(Numeric); ok{
			return f.Numeric()<d.Numeric()
		} else {
			var message = fmt.Sprintf("%v can't compare with %v", f, o)
			panic(message)
		}
	}
}
func(f Float)Compare(other interface{}) int {
	switch o := other.(type) {
	case int:
		return CompareFF(f.Numeric(), float64(o))
	case float64:
		return CompareFF(f.Numeric(), o)
	default:
		if d, ok :=o.(Numeric); ok{
			return CompareFF(f.Numeric(), d.Numeric())
		} else {
			var message = fmt.Sprintf("%v can't compare with %v", f, o)
			panic(message)
		}
	}
}
func CompareFF(x float64, y float64) int {
	if x>y {
		return 1
	}
	if x<y {
		return -1
	}
	return 0
}


type Stringer interface {
	String() string
}

type String string
func (s String)String() string{
	return string(s)
}
func (s String)Less(other interface{}) bool{
	switch o := other.(type) {
	case string:
		return s.String()<o
	case String:
		return s<o
	default:
		if d, ok := o.(Stringer); ok{
			return s.String()<d.String()
		}else{
			var message = fmt.Sprintf("%v can't compare with %v", s, o)
			panic(message)
		}
	}
}
func(s String)Compare(other interface{}) int {
	switch o := other.(type) {
	case string:
		return CompareSS(string(s), o)
	case String:
		return CompareSS(string(s), string(o))
	default:
		if d, ok :=o.(Stringer); ok{
			return CompareSS(s.String(), d.String())
		} else {
			var message = fmt.Sprintf("%v can't compare with %v", s, o)
			panic(message)
		}
	}
}

func CompareSS(x string, y string) int {
	if x>y {
		return 1
	}
	if x<y {
		return -1
	}
	return 0
}

type Sortable []Compare

func (s Sortable)Len() int {
	return len(s)
}

func (s Sortable)Less(i, j int) bool {
	return s[i].Less(s[j])
}

func (s Sortable)Swap(i, j int){
	s[i], s[j] = s[j], s[i]
}


