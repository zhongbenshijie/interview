package defer_test

import (
	"fmt"
	"testing"
)

func defer_call() {
	defer func() {
		fmt.Println("11111")
	}()
	defer func() {
		fmt.Println("22222")
	}()

	// 如果没有recover, panic将最后输出
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

func TestDefer(t *testing.T) {
	defer_call()
}
