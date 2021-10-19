package check_element_is_in_slice

import (
	"reflect"
	"testing"
)

// 检查某个元素是否在切片中
func CheckElementIsInSlice(slice interface{}, element interface{}) bool {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		return false
	}

	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == element {
			return true
		}
	}

	return false
}

func TestCheckElementIsInSlice(t *testing.T) {
	strSlice := []string{"India", "Canada", "Japan"}
	t.Log(CheckElementIsInSlice(strSlice, "India"))
	t.Log(CheckElementIsInSlice(strSlice, "American"))

	// check_element_is_in_slice_test.go:27: true
	//    check_element_is_in_slice_test.go:28: false
}
