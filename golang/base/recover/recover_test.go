package recover_test

import (
	"errors"
	"testing"
)

func TestRecover(t *testing.T) {
	// 这段代码是比较危险的，因为要当心recover在做的事情，因为我们不对recover的错误进行判断
	// 这样如果是系统的一些核心资源消耗完了，那么就会报错，而我们又不让程序挂掉，也不进行报错检查
	// 只是简单的做了一下错误记录，这样程序虽然还在跑，但是无法提供服务了，就相当于一个僵尸进程
	// 所以不如不进行recover，而是让程序挂掉，这样当程序挂掉，我们的守护进程就会检测到，就会帮我们重启程序
	defer func() {
		if err := recover(); err != nil {
			t.Log("recovered from", err)
		}
	}()

	panic(errors.New("Something wrong!"))
}
