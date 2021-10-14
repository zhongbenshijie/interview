package variable_params_test

import (
	"fmt"
	"testing"
)

// 可变参数的意思就是可以不指定参数的个数，但是参数的类型都是一样的
func variableParams(opt ...int) {
	for _, op := range opt {
		fmt.Println(op)
	}
}

func TestVariableParams(t *testing.T) {
	variableParams(1, 2)
	variableParams(1, 2, 3)
}
