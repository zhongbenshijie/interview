package run_once

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {
}

var (
	singletonInstance *Singleton
	once              sync.Once
)

func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("Create Obj")
		singletonInstance = new(Singleton)
	})

	return singletonInstance
}

func TestGetSingletonObj(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			obj := GetSingletonObj()
			fmt.Printf("%d\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}

	wg.Wait()

	// 	Create Obj
	// 19185592
	// 19185592
	// 19185592
	// 19185592
	// 19185592
	// 19185592
	// 19185592
	// 19185592
	// 19185592
	// 19185592

	// 只运行一次的意思是，在并发的环境下，在一个进程中，包在once.Do里面的这个函数只会被执行一次
}
