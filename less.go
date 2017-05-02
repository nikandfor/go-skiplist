package skiplist

var (
	IntLess LessFunc = func(a, b interface{}) bool {
		return a.(int) < b.(int)
	}
	IntGreater LessFunc = func(a, b interface{}) bool {
		return a.(int) > b.(int)
	}
	IntLE LessFunc = func(a, b interface{}) bool {
		return a.(int) <= b.(int)
	}
	IntGE LessFunc = func(a, b interface{}) bool {
		return a.(int) >= b.(int)
	}
	Int64Less LessFunc = func(a, b interface{}) bool {
		return a.(int64) < b.(int64)
	}
	Int64Greater LessFunc = func(a, b interface{}) bool {
		return a.(int64) > b.(int64)
	}
	Uint64Less LessFunc = func(a, b interface{}) bool {
		return a.(uint64) < b.(uint64)
	}
	Uint64Greater LessFunc = func(a, b interface{}) bool {
		return a.(uint64) > b.(uint64)
	}
	Int64LE LessFunc = func(a, b interface{}) bool {
		return a.(int64) <= b.(int64)
	}
	Int64GE LessFunc = func(a, b interface{}) bool {
		return a.(int64) >= b.(int64)
	}
	Uint64LE LessFunc = func(a, b interface{}) bool {
		return a.(uint64) <= b.(uint64)
	}
	Uint64GE LessFunc = func(a, b interface{}) bool {
		return a.(uint64) >= b.(uint64)
	}
	Int32Less LessFunc = func(a, b interface{}) bool {
		return a.(int32) < b.(int32)
	}
	Int32Greater LessFunc = func(a, b interface{}) bool {
		return a.(int32) > b.(int32)
	}
	Uint32Less LessFunc = func(a, b interface{}) bool {
		return a.(uint32) < b.(uint32)
	}
	Uint32Greater LessFunc = func(a, b interface{}) bool {
		return a.(uint32) > b.(uint32)
	}
	Int32LE LessFunc = func(a, b interface{}) bool {
		return a.(int32) <= b.(int32)
	}
	Int32GE LessFunc = func(a, b interface{}) bool {
		return a.(int32) >= b.(int32)
	}
	Uint32LE LessFunc = func(a, b interface{}) bool {
		return a.(uint32) <= b.(uint32)
	}
	Uint32GE LessFunc = func(a, b interface{}) bool {
		return a.(uint32) >= b.(uint32)
	}
	StringLess LessFunc = func(a, b interface{}) bool {
		return a.(string) < b.(string)
	}
	StringGreater LessFunc = func(a, b interface{}) bool {
		return a.(string) > b.(string)
	}
	StringLE LessFunc = func(a, b interface{}) bool {
		return a.(string) <= b.(string)
	}
	StringGE LessFunc = func(a, b interface{}) bool {
		return a.(string) >= b.(string)
	}
)
