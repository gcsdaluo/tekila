package slice

import (
	"github.com/gcsdaluo/tekila/internal/common"
	"github.com/gcsdaluo/tekila/internal/err"
)

// RemoveAtIndex 删除指定索引的元素:返回新切片，删除元素，错误处理
func RemoveAtIndex[T any](src []T, index int) ([]T, T, error) {
	// 错误处理
	if index < 0 || index >= len(src) {
		return nil, common.ZeroValueOfT[T](), err.ErrIndexOutOfRange(len(src), index)
	}

	// 取出index下的元素,使用copy原地删除操作
	result := src[index]
	copy(src[index:], src[index+1:])

	return src[:len(src)-1], result, nil
}

// RemoveValue 删除指定值的元素(首次出现)
func RemoveValue[T comparable](src []T, target T) ([]T, T, error) {
	for i, elem := range src {
		if elem == target {
			return append(src[:i], src[i+1:]...), target, nil
		}
	}

	return src, target, err.NewError("Could not find the specified value")
}

// RemoveAllValue 删除指定值元素(所有指定的元素)
func RemoveAllValue[T comparable](src []T, target T) ([]T, T, error) {
	// 尝试预估容量，尽量使用src或者src一半，因为删除元素一般都是一部分的符合条件的
	result := make([]T, 0, len(src)/2)

	for _, elem := range src {
		if elem != target {
			result = append(result, elem)
		}
	}

	return result, target, err.NewError("Could not find the specified value")
}

// FilterRemove 删除满足条件的元素
func FilterRemove[T any](src []T, condition func(index int, src T) bool) []T {
	// 记录被删除元素位置
	emptyPos := 0
	for index := range src {
		// 判断是否满足condition
		if condition(index, src[index]) {
			continue
		}

		// 不满足条件，移动元素
		src[emptyPos] = src[index]
		emptyPos++
	}

	// 最终返回不包括emptyPos的slice
	return src[:emptyPos]
}

// RemoveDuplicates 删除指定重复的元素，仅保留一个副本
func RemoveDuplicates[T comparable](src []T, target T) ([]T, error) {
	if len(src) == 0 {
		return nil, err.NewError("The slice are empty")
	}

	// 设置标志位和新的切片数组
	result := make([]T, 0, len(src)/2)
	found := false

	// 循环遍历切片寻找重复的指定元素，在重复元素加入切片后标志为设为true
	for _, elem := range src {
		if elem == target && !found {
			found = true
			result = append(result, elem)
		} else if elem != target {
			result = append(result, elem)
		}
	}

	return result, nil
}

// RemoveAllDuplicates 删除所有重复的元素，仅保留一个副本
func RemoveAllDuplicates[T comparable](src []T) ([]T, error) {
	if len(src) == 0 {
		return nil, err.NewError("The slice are empty")
	}

	// 设置每个重复元素对应的map标志为和切片数组
	mapping := map[T]bool{}
	result := make([]T, 0, len(src)/2)

	for _, elem := range src {
		if !mapping[elem] {
			mapping[elem] = true
			result = append(result, elem)
		}
	}

	return result, nil
}

// RemoveRangeAtIndex 删除位于给定索引范围内的元素
func RemoveRangeAtIndex[T any](src []T, start, end int) ([]T, []T, error) {
	if start < 0 || end >= len(src) || start > end {
		return nil, nil, err.NewError("Invalid index range")
	}

	removed := src[start : end+1]
	copy(src[start:], src[end+1:])
	return src[:len(src)-(end-start+1)], removed, nil
}
