package sort

import(
	"reflect"
)

func IntLess(x, y int)bool{
	return x<y
}
var vIntLess = reflect.ValueOf(IntLess)
func Int8Less(x, y int8)bool{
	return x<y
}

func Int16Less(x, y int16)bool{
	return x<y
}
func Int32Less(x, y int32)bool{
	return x<y
}
func Int64Less(x, y int64)bool{
	return x<y
}
func UintLess(x, y uint)bool{
	return x<y
}
func Uint8Less(x, y uint8)bool{
	return x<y
}
func Uint16Less(x, y uint16)bool{
	return x<y
}
func Uint32Less(x, y uint32)bool{
	return x<y
}
func Uint64Less(x, y uint64)bool{
	return x<y
}
func Float32Less(x, y float32)bool{
	return x<y
}
func Float64Less(x, y float64)bool{
	return x<y
}
func RuneLess(x, y rune)bool{
	return x<y
}
func StringLess(x, y string)bool{
	return x<y
}

func IntCompare(x, y int)int{
	if x<y {
		return -1
	}
	if x>y {
		return 1
	}
	return 0
}
func Int8Compare(x, y int8)int{
	if x<y {
		return -1
	}
	if x>y {
		return 1
	}
	return 0
}
func Int16Compare(x, y int16)int{
	if x<y {
		return -1
	}
	if x>y {
		return 1
	}
	return 0
}
func Int32Compare(x, y int32)int{
	if x<y {
		return -1
	}
	if x>y {
		return 1
	}
	return 0
}
func Int64Compare(x, y int64)int{
	if x<y {
		return -1
	}
	if x>y {
		return 1
	}
	return 0
}
func UintCompare(x, y uint)int{
	if x<y {
		return -1
	}
	if x>y {
		return 1
	}
	return 0
}
func Uint8Compare(x, y uint8)int{
	if x<y {
		return -1
	}
	if x>y {
		return 1
	}
	return 0
}
func Uint16Compare(x, y uint16)int{
	if x<y {
		return -1
	}
	if x>y {
		return 1
	}
	return 0
}
func Uint32Compare(x, y uint32)int{
	if x<y {
		return -1
	}
	if x>y {
		return 1
	}
	return 0
}
func Uint64Compare(x, y uint64)int{
	if x<y {
		return -1
	}
	if x>y {
		return 1
	}
	return 0
}
func Float32Compare(x, y float32)int{
	if x<y {
		return -1
	}
	if x>y {
		return 1
	}
	return 0
}
func Float64Compare(x, y float64)int{
	if x<y {
		return -1
	}
	if x>y {
		return 1
	}
	return 0
}
func RuneCompare(x, y rune)int{
	if x<y {
		return -1
	}
	if x>y {
		return 1
	}
	return 0
}
func StringCompare(x, y string)int{
	if x<y {
		return -1
	}
	if x>y {
		return 1
	}
	return 0
}

var IntVCompare = func(args []reflect.Value)[]reflect.Value{
	var ret = []reflect.Value{reflect.ValueOf(
		Int64Compare(args[0].Int(), args[1].Int()))}
	return ret
}
var UintVCompare = func(args []reflect.Value)[]reflect.Value{
	var ret = []reflect.Value{reflect.ValueOf(
		Uint64Compare(args[0].Uint(), args[1].Uint()))}
	return ret
}
var FloatVCompare = func(args []reflect.Value)[]reflect.Value{
	var ret = []reflect.Value{reflect.ValueOf(
		Float64Compare(args[0].Float(), args[1].Float()))}
	return ret
}
var StringVCompare = func(args []reflect.Value)[]reflect.Value{
	var ret = []reflect.Value{reflect.ValueOf(
		StringCompare(args[0].String(), args[1].String()))}
	return ret
}

var IntVLess = func(args []reflect.Value)[]reflect.Value{
	var ret = []reflect.Value{reflect.ValueOf(
		Int64Less(args[0].Int(), args[1].Int()))}
	return ret
}
var UintVLess = func(args []reflect.Value)[]reflect.Value{
	var ret = []reflect.Value{reflect.ValueOf(
		Uint64Less(args[0].Uint(), args[1].Uint()))}
	return ret
}
var FloatVLess = func(args []reflect.Value)[]reflect.Value{
	var ret = []reflect.Value{reflect.ValueOf(
		Float64Less(args[0].Float(), args[1].Float()))}
	return ret
}
var StringVLess = func(args []reflect.Value)[]reflect.Value{
	var ret = []reflect.Value{reflect.ValueOf(
		StringLess(args[0].String(), args[1].String()))}
	return ret
}

var registedLess = map[reflect.Kind]reflect.Value{
	reflect.Int:reflect.ValueOf(IntLess),
	reflect.Int8:reflect.ValueOf(Int8Less),
	reflect.Int16:reflect.ValueOf(Int16Less),
	reflect.Int32:reflect.ValueOf(Int32Less),
	reflect.Int64:reflect.ValueOf(Int64Less),
	reflect.Uint:reflect.ValueOf(UintLess),
	reflect.Uint8:reflect.ValueOf(Uint8Less),
	reflect.Uint16:reflect.ValueOf(Uint16Less),
	reflect.Uint32:reflect.ValueOf(Uint32Less),
	reflect.Uint64:reflect.ValueOf(Uint64Less),
	reflect.Float32:reflect.ValueOf(Float32Less),
	reflect.Float64:reflect.ValueOf(Float64Less),
	reflect.String:reflect.ValueOf(StringLess),
}
var registedCompare = map[reflect.Kind]reflect.Value{
	reflect.Int:reflect.ValueOf(IntCompare),
	reflect.Int8:reflect.ValueOf(Int8Compare),
	reflect.Int16:reflect.ValueOf(Int16Compare),
	reflect.Int32:reflect.ValueOf(Int32Compare),
	reflect.Int64:reflect.ValueOf(Int64Compare),
	reflect.Uint:reflect.ValueOf(UintCompare),
	reflect.Uint8:reflect.ValueOf(Uint8Compare),
	reflect.Uint16:reflect.ValueOf(Uint16Compare),
	reflect.Uint32:reflect.ValueOf(Uint32Compare),
	reflect.Uint64:reflect.ValueOf(Uint64Compare),
	reflect.Float32:reflect.ValueOf(Float32Compare),
	reflect.Float64:reflect.ValueOf(Float64Compare),
	reflect.String:reflect.ValueOf(StringCompare),
}

func MakeLess(fptr interface{}) {
	var fval = reflect.ValueOf(fptr).Elem()
	var ftype = fval.Type()
	var less = LessMaker(ftype)
	fval.Set(less)
}

func MakeCompare(fptr interface{}) {
	var fval = reflect.ValueOf(fptr).Elem()
	var ftype = fval.Type()
	var compare = CompareMaker(ftype)
	fval.Set(compare)
}

func MakeStructCompare(fptr interface{}, by... string){
	var fval = reflect.ValueOf(fptr).Elem()
	var ftype = fval.Type()
	var compare = func(args []reflect.Value)[]reflect.Value{
		var argx = args[0]
		var argy = args[1]
		for _, fname := range by {
			var fieldx = argx.FieldByName(fname)
			var fieldy = argy.FieldByName(fname)
			var fieldType = fieldx.Type()
			var rec []reflect.Value
			switch fieldType.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				rec = IntVCompare([]reflect.Value{fieldx, fieldy})
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				rec = UintVCompare([]reflect.Value{fieldx, fieldy})
			case reflect.Float64, reflect.Float32:
				rec = FloatVCompare([]reflect.Value{fieldx, fieldy})
			case reflect.String:
				rec = StringVCompare([]reflect.Value{fieldx, fieldy})
			default:
				panic("I don't know how to compare it")
			}
			if rec[0].Int() != 0 {
				return rec
			}
		}
		return []reflect.Value{reflect.ValueOf(0)}
	}
	fval.Set(reflect.MakeFunc(ftype, compare))
}

func MakeStructLess(fptr interface{}, by... string){
	var fval = reflect.ValueOf(fptr).Elem()
	var ftype = fval.Type()
	var compare = func(args []reflect.Value)[]reflect.Value{
		var argx = args[0]
		var argy = args[1]
		for _, fname := range by {
			var fieldx = argx.FieldByName(fname)
			var fieldy = argy.FieldByName(fname)
			var fieldType = fieldx.Type()
			var rec []reflect.Value
			switch fieldType.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				rec = IntVCompare([]reflect.Value{fieldx, fieldy})
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				rec = UintVCompare([]reflect.Value{fieldx, fieldy})
			case reflect.Float64, reflect.Float32:
				rec = FloatVCompare([]reflect.Value{fieldx, fieldy})
			case reflect.String:
				rec = StringVCompare([]reflect.Value{fieldx, fieldy})
			default:
				panic("I don't know how to compare it")
			}
			var ret = rec[0].Int()
			if ret != 0 {
				return []reflect.Value{reflect.ValueOf(ret==-1)}
			}
		}
		return []reflect.Value{reflect.ValueOf(false)}
	}
	fval.Set(reflect.MakeFunc(ftype, compare))
}

func LessMaker(ftype reflect.Type)reflect.Value{
	var in1st = ftype.In(0).Kind()
	switch in1st {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return reflect.MakeFunc(ftype, IntVLess)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return reflect.MakeFunc(ftype, UintVLess)
	case reflect.Float64, reflect.Float32:
		return reflect.MakeFunc(ftype, FloatVLess)
	case reflect.String:
		return reflect.MakeFunc(ftype, StringVLess)
	default:
		panic("I don't know how to compare it")
	}
}

func CompareMaker(ftype reflect.Type)reflect.Value{
	var in1st = ftype.In(0).Kind()
	switch in1st {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return reflect.MakeFunc(ftype, IntVCompare)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return reflect.MakeFunc(ftype, UintVCompare)
	case reflect.Float64, reflect.Float32:
		return reflect.MakeFunc(ftype, FloatVCompare)
	case reflect.String:
		return reflect.MakeFunc(ftype, StringVCompare)
	default:
		panic("I don't know how to compare it")
	}
}
