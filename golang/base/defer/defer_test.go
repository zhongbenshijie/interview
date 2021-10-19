package defer1

import (
	"fmt"
	"testing"
)

// defer后进先出的执行
// 如果没有recover来捕捉panic，panic将最后输出
// 如果有recover来捕捉panic，那么panic将在recover位置输出，不在最后输出

func deferCall1() {
	defer func() {
		fmt.Println("11111")
	}()
	defer func() {
		fmt.Println("22222")
	}()

	// 如果没有recover, panic将最后输出!!!
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover from r : ", r)
		}
	}()

	defer func() {
		fmt.Println("33333")
	}()

	fmt.Println("111 Helloworld")

	panic("Panic 1!")

	panic("Panic 2!")

	fmt.Println("222 Helloworld")
}

func TestDefer1(t *testing.T) {
	deferCall1()

	// 111 Helloworld
	// 33333
	// Recover from r :  Panic 1!
	// 22222
	// 11111
}

func deferCall2() (i int) {
	defer func() { i++ }()

	return 1
}

func TestDefer2(t *testing.T) {
	t.Log(deferCall2())

	// 2
	// 返回值被提前声明，defer在return后执行，所以能够改变返回值
}

func deferCall3() int {
	var i int
	defer func() {
		i++
		fmt.Println("deferCall3:", i)
	}()
	defer func() {
		i++
		fmt.Println("deferCall3:", i)
	}()

	return i
}

func TestDefer3(t *testing.T) {
	t.Log(deferCall3())

	// 1
	// 2
	// 0
	// 返回值没有被提前声明，defer在return后执行，返回值没有被改写
}

func deferCall4() *int {
	var i int
	defer func() {
		i++
		fmt.Println("deferCall4:", i) // 打印结果为 c defer: 2
	}()
	defer func() {
		i++
		fmt.Println("deferCall4:", i) // 打印结果为 c defer: 1
	}()
	return &i
}

func TestDefer4(t *testing.T) {
	t.Log(*deferCall4())

	// 1
	// 2
	// 2
	// 虽然返回值没有被提前声明，但是返回值是指针类型，所以返回值被改写了
}
