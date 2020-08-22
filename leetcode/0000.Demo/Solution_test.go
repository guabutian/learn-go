package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs bool
		expect bool
	}{
		{"TestCase", true, true},
		{"TestCase", true, true},
		{"TestCase", false, false},
	}

	//	开始测试
	for i, c := range cases {
		
		// 前面这个是case，后面是它的代码
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			// 输入参数，并获得返回值
			got := Solution(c.inputs)
			
			// 深度对比
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
			}
		})
	}
}

//	压力测试
func BenchmarkSolution(b *testing.B) {

}

//	使用案列
func ExampleSolution() {

}
