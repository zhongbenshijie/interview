package const1

import (
	"testing"
)

const (
	Monday = 1 + iota
	Tuesday
	Wendesday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTry(t *testing.T) {
	// t.Log(iota) crash error: undefined: iota
	t.Log(Monday, Tuesday, Wendesday)
}

func TestConstantTry1(t *testing.T) {
	t.Log(Readable, Writable, Executable) // 1 2 4
	a := 1                                // 0001

	// 可以很好的做状态位的标识
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
