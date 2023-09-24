package slice

// UnionSet 切片并集 返回元素顺序不固定
func UnionSet[T comparable](src, dst []T) []T {
	srcMapping, dstMapping := sToMap[T](src), sToMap[T](dst)
	for key := range srcMapping {
		dstMapping[key] = struct{}{}
	}

	var res = make([]T, 0, len(dstMapping))
	for key := range srcMapping {
		res = append(res, key)
	}

	return res
}

// UnionSetFunc 切片并集
// 通过自定义比较函数
func UnionSetFunc[T any](src, dst []T, equal equalFunc[T]) {
	var target = make([]T, 0, len(src)+len(dst))
	target = append(target, src...)
	target = append(target, dst...)

	return
}
