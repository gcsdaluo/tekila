package slice

import (
	"testing"
)

func TestRemoveAtIndex(t *testing.T) {
	// 定义测试用例

	// Define a custom Person struct for demonstration purposes
	type Person struct {
		Name string
	}

	testCases := []struct {
		numberSlice     []int
		index           int
		expected        []int
		expectedRemoved int
		expectErr       bool
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 4, 5}, 3, false},
		{[]int{10, 20, 30, 40}, 0, []int{20, 30, 40}, 10, false},
		{[]int{100, 200, 300}, 2, []int{100, 200}, 300, false},
		{[]int{1, 2, 3}, -1, nil, 0, true},
		{[]int{1, 2, 3}, 3, nil, 0, true},
	}

	// 执行测试用例
	for _, tc := range testCases {
		newSlice, removed, err := RemoveAtIndex(tc.numberSlice, tc.index)

		// 验证是否有错误发生
		if (err != nil) != tc.expectErr {
			t.Errorf("Expected error: %v, but got: %v", tc.expectErr, err)
		}

		// 验证切片是否与预期结果一致
		if !equalSlices(newSlice, tc.expected) {
			t.Errorf("Expected: %v, but got: %v", tc.expected, newSlice)
		}

		// 验证被删除的元素是否与预期结果一致
		if removed != tc.expectedRemoved {
			t.Errorf("Expected removed element: %v, but got: %v", tc.expectedRemoved, removed)
		}

		//log.Println("New Slice:", newSlice)
		//log.Println("Removed Element:", removed)

	}
}

// helper函数，用于比较两个切片是否相等
func equalSlices(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}
	return true
}
