package array_and_slice

import (
	"testing"
)

// 数组是一种带长度的类型，长度不同也意味着是不同的类型
// 切片是长度可变的类数组存储结构，切片本质是一个结构体，里面有一个成员指针指向一个数组，这个数组的长度是可以自增长的
// 切片这个结构体有三个成员，ptr，len，cap，ptr指向一片连续的存储空间

// 切片是如何进行自增长的
func TestSliceGrowing(t *testing.T) {
	s := []int{}

	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}

	// array_and_slice_test.go:17: 1 1
	// array_and_slice_test.go:17: 2 2
	// array_and_slice_test.go:17: 3 4
	// array_and_slice_test.go:17: 4 4
	// array_and_slice_test.go:17: 5 8
	// array_and_slice_test.go:17: 6 8
	// array_and_slice_test.go:17: 7 8
	// array_and_slice_test.go:17: 8 8
	// array_and_slice_test.go:17: 9 16
	// array_and_slice_test.go:17: 10 16

	// 根据输出结果可以看到，cap的增长是当不够用时(len > cap, len等于cap时不会进行自增长)就会翻倍申请空间
}

// 切片共享存储空间
func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Otc", "Nov", "Dec"}

	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))

	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "Unknow"
	t.Log(Q2)
	t.Log(year)

	// array_and_slice_test.go:39: [Apr May Jun] 3 9
	// array_and_slice_test.go:42: [Jun Jul Aug] 3 7
	// array_and_slice_test.go:44: [Apr May Unknow]
	// array_and_slice_test.go:45: [Jan Feb Mar Apr May Unknow Jul Aug Sep Otc Nov Dec]

	// 根据输出结果可以看到，截取一个切片，那么直至切片结束的那部分存储空间都是可以访问修改的，cap取值时就是根据这个来计算的
	// 所以当一个切片被截取修改时，可能会对其他截取的结果造成影响，因为内存本质是同一块
}
