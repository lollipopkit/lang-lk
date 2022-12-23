package api

type GoFunction func(LkState) int

func LkUpvalueIndex(i int) int {
	return LK_REGISTRYINDEX - i
}

type LkState interface {
	BasicAPI
	AuxLib
}

type BasicAPI interface {
	/* basic stack manipulation */
	GetTop() int
	AbsIndex(idx int) int
	CheckStack(n int) bool
	Pop(n int)
	Copy(fromIdx, toIdx int)
	PushValue(idx int)
	Replace(idx int)
	Insert(idx int)
	Remove(idx int)
	Rotate(idx, n int)
	SetTop(idx int)
	XMove(to LkState, n int)
	/* access functions (stack -> Go) */
	TypeName(tp LkType) string
	Type(idx int) LkType
	IsNone(idx int) bool
	IsNil(idx int) bool
	IsNoneOrNil(idx int) bool
	IsBoolean(idx int) bool
	IsInteger(idx int) bool
	IsNumber(idx int) bool
	IsString(idx int) bool
	IsTable(idx int) bool
	IsThread(idx int) bool
	IsFunction(idx int) bool
	IsGoFunction(idx int) bool
	ToBoolean(idx int) bool
	ToInteger(idx int) int64
	ToIntegerX(idx int) (int64, bool)
	ToNumber(idx int) float64
	ToNumberX(idx int) (float64, bool)
	ToString(idx int) string
	ToStringX(idx int) (string, bool)
	ToGoFunction(idx int) GoFunction
	ToThread(idx int) LkState
	ToPointer(idx int) interface{}
	/* push functions (Go -> stack) */
	PushNil()
	PushBoolean(b bool)
	PushInteger(n int64)
	PushNumber(n float64)
	PushString(s string)
	PushFString(fmt string, a ...interface{})
	PushGoFunction(f GoFunction)
	PushGoClosure(f GoFunction, n int)
	PushGlobalTable()
	PushThread() bool
	/* Comparison and arithmetic functions */
	Arith(op ArithOp)
	Compare(idx1, idx2 int, op CompareOp) bool
	/* get functions (Lua -> stack) */
	NewTable()
	CreateTable(nArr, nRec int)
	GetTable(idx int) LkType
	GetField(idx int, k string) LkType
	GetI(idx int, i int64) LkType
	RawGet(idx int) LkType
	RawGetI(idx int, i int64) LkType
	GetGlobal(name string) LkType
	/* set functions (stack -> Lua) */
	SetTable(idx int)
	SetField(idx int, k string)
	SetMetatable(idx int)
	SetI(idx int, i int64)
	RawSet(idx int)
	RawSetI(idx int, i int64)
	SetGlobal(name string)
	Register(name string, f GoFunction)
	/* 'load' and 'call' functions (load and run Lua code) */
	Load(chunk []byte, chunkName, mode string) LkStatus
	Call(nArgs, nResults int)
	PCall(nArgs, nResults, msgh int) LkStatus
	/* miscellaneous functions */
	Len(idx int)
	Next(idx int) bool
	Error() int
	StringToNumber(s string) bool
	/* coroutine functions */
	NewThread() LkState
	Resume(from LkState, nArgs int) LkStatus
	Yield(nResults int) LkStatus
	Status() LkStatus
	IsYieldable() bool
	GetStack() bool // debug

	CatchAndPrint()
}