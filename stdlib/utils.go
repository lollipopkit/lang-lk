package stdlib

import (
	"fmt"
	"reflect"

	. "git.lolli.tech/lollipopkit/go-lang-lk/api"
)

func pushValue(ls LkState, item any) {
	switch i := item.(type) {
	case string:
		ls.PushString(i)
	case int64:
		ls.PushInteger(i)
	case float64:
		ls.PushNumber(i)
	case bool:
		ls.PushBoolean(i)
	case GoFunction:
		ls.PushGoFunction(i)
	case nil:
		ls.PushNil()
	default:
		v := reflect.ValueOf(i)
		switch v.Kind() {
		case reflect.Slice:
			items := make([]any, v.Len())
			for i := 0; i < v.Len(); i++ {
				items[i] = v.Index(i).Interface()
			}
			pushList(ls, items)
			return
		case reflect.Map:
			items := make(map[string]any)
			for _, key := range v.MapKeys() {
				items[key.String()] = v.MapIndex(key).Interface()
			}
			pushTable(ls, items)
			return
		}
		panic(fmt.Sprintf("unsupported type: %T", item))
	}
}

func pushList(ls LkState, items []any) {
	ls.CreateTable(len(items), 0)
	for i, item := range items {
		pushValue(ls, item)
		ls.SetI(-2, int64(i+1))
	}
}

func pushTable(ls LkState, items map[string]any) {
	ls.CreateTable(0, len(items)+1)
	for k, v := range items {
		pushValue(ls, v)
		ls.SetField(-2, k)
	}
}

func getTable(ls LkState, idx int) map[string]any {
	ls.CheckType(idx, LUA_TTABLE)
	table := make(map[string]any)
	ls.PushNil()
	for ls.Next(idx) {
		key := ls.ToString(-2)
		val := ls.ToPointer(-1)
		table[key] = val
		ls.Pop(1)
	}
	return table
}

func OptTable(ls LkState, idx int, dft map[string]any) map[string]any {
	if ls.IsNoneOrNil(idx) {
		return dft
	}
	return getTable(ls, idx)
}

// lua-5.3.4/src/loslib.c#getfield()
func _getField(ls LkState, key string, dft int64) int {
	t := ls.GetField(-1, key) /* get field and its type */
	res, isNum := ls.ToIntegerX(-1)
	if !isNum { /* field is not an integer? */
		if t != LUA_TNIL { /* some other value? */
			return ls.Error2("field '%s' is not an integer", key)
		} else if dft < 0 { /* absent field; no default? */
			return ls.Error2("field '%s' missing in date table", key)
		}
		res = dft
	}
	ls.Pop(1)
	return int(res)
}

func getFunc(ls LkState, idx int) GoFunction {
	ls.CheckType(idx, LUA_TFUNCTION)
	return ls.ToGoFunction(idx)
}
