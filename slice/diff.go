package slice

// DiffSet 差集，只支持 comparable 类型
// 已去重
// 并且返回值的顺序是不确定的
func DiffSet[T comparable](src, target []T) []T {
	srcMap := sToMap[T](src)
	for _, elem := range target {
		delete(srcMap, elem)
	}

	var res = make([]T, 0, len(srcMap))
	for key := range srcMap {
		res = append(res, key)
	}
	return res
}

// DiffSetFunc 差集，已去重
// 你应该优先使用 DiffSet
func DiffSetFunc[T any](src, target []T, equal equalFunc[T]) []T {
	// TODO 优化容量预估
	var res = make([]T, 0, len(src))
	for _, elem := range src {
		if !ContainsFunc[T](target, elem, equal) {
			res = append(res, elem)
		}
	}
	return deduplicateFunc[T](res, equal)
}
