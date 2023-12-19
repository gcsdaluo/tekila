package slice

// Find 在切片中查找符合条件的元素
func Find[T any](src []T, condition conditionFunc[T]) (T, bool) {
	for _, elem := range src {
		if condition(elem) {
			return elem, true
		}
	}

	// 如果没有找到符合条件的元素，返回零值和 false
	var zeroValue T
	return zeroValue, false
}

// FindAll 查找所有符合条件的元素
// 不会返回 nil
func FindAll[T any](src []T, condition conditionFunc[T]) []T {
	// 结果切片需要做成 1/8 即触发扩容的情况下，最多三次就会和原本的容量一样
	// +1 是为了保证，至少有一个元素
	res := make([]T, 0, len(src)>>3+1)
	for _, elem := range src {
		if condition(elem) {
			res = append(res, elem)
		}
	}
	return res
}
