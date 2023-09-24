package slice

// FirstIndex 0返回与target相等的第一个元素下标
// -1表示未找到
func FirstIndex[T comparable](src []T, target T) int {
	return FirstIndexFunc[T](src, target, func(src, dst T) bool {
		return src == dst
	})
}

// FirstIndexFunc 返回与target相等的第一个元素下标
// 通过自定义函数获取
// -1表示未找到
func FirstIndexFunc[T comparable](src []T, target T, equal equalFunc[T]) int {
	for i, elem := range src {
		if equal(elem, target) {
			return i
		}
	}
	return -1
}

// LastIndex 返回和target相等的最后一个元素的下标
// -1表示未找到
func LastIndex[T comparable](src []T, target T) int {
	return LastIndexFunc[T](src, target, func(src, dst T) bool {
		return src == dst
	})
}

// LastIndexFunc 返回和target相等的最后一个元素的下标
// 通过自定义函数获取
// -1表示未找到
func LastIndexFunc[T any](src []T, target T, equal equalFunc[T]) int {
	for i := len(src) - 1; i >= 0; i-- {
		if equal(src[i], target) {
			return i
		}
	}
	return -1
}

// IndexAll 返回和 target元素相等的所有元素的下标
func IndexAll[T comparable](src []T, target T) []int {
	return IndexAllFunc[T](src, target, func(src, dst T) bool {
		return src == dst
	})
}

// IndexAllFunc 返回和 target元素相等的所有元素的下标
// 通过自定义函数获取
func IndexAllFunc[T any](src []T, target T, equal equalFunc[T]) []int {
	var indexList = make([]int, 0, len(src))
	for i, elem := range src {
		if equal(elem, target) {
			indexList = append(indexList, i)
		}
	}
	return indexList
}
