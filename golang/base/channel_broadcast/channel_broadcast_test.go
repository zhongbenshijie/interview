package channel_broadcast

import (
	"testing"
	"time"
)

func isCanceled(cancelCh chan struct{}) bool {
	select {
	case <-cancelCh:
		return true
	default:
		return false
	}
}

func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{}, 0)

	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCanceled(cancelCh) {
					break
				}

				time.Sleep(time.Millisecond * 5)
			}
			t.Log(i, "Canneled")
		}(i, cancelChan)
	}

	// 关闭channel后，channel不会阻塞，而是会立即返回，这样就形成了一种类广播式的信号传递
	// 任何监听这个channel是否关闭的协程都可以依赖于这个信号进行逻辑处理
	close(cancelChan)
	time.Sleep(time.Second)

	// channel_broadcast_test.go:29: 4 Canneled
	// channel_broadcast_test.go:29: 0 Canneled
	// channel_broadcast_test.go:29: 3 Canneled
	// channel_broadcast_test.go:29: 1 Canneled
	// channel_broadcast_test.go:29: 2 Canneled
}
