package slice

// SymmetricDiffSet 对称差集
// 已去重
// 返回值的元素顺序是不定的
func SymmetricDiffSet[T comparable](src, dst []T) []T {
	srcMap, dstMap := sToMap[T](src), sToMap[T](dst)
	for key := range dstMap {
		if _, ok := srcMap[key]; ok {
			delete(srcMap, key)
		} else {
			srcMap[key] = struct{}{}
		}
	}

	res := make([]T, 0, len(srcMap))
	for key := range srcMap {
		res = append(res, key)
	}

	return res
}

// SymmetricDiffSetFunc 对称差集
// 已去重
func SymmetricDiffSetFunc[T any](src, dst []T, equal equalFunc[T]) []T {
	res := []T{}

	// 找出在 src 不在 dst 的元素
	for _, elem := range src {
		if !ContainsFunc[T](dst, func(target T) bool {
			return equal(target, elem)
		}) {
			res = append(res, elem)
		}
	}

	// 找出在 dst 不在 src 的元素
	for _, elem := range dst {
		if !ContainsFunc[T](src, func(target T) bool {
			return equal(target, elem)
		}) {
			res = append(res, elem)
		}
	}

	return deduplicateFunc[T](res, equal)
}
