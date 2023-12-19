package slice

// IntersectSet 取交集，只支持 comparable 类型
// 已去重
func IntersectSet[T comparable](src []T, dst []T) []T {
	srcMap := sToMap(src)
	var ret = make([]T, 0, len(src))

	// 交集小于等于两个集合中的任何一个
	for _, elem := range dst {
		if _, exist := srcMap[elem]; exist {
			ret = append(ret, elem)
		}
	}

	return deduplicate[T](ret)
}

// IntersectSetFunc 支持任意类型
// 已去重
func IntersectSetFunc[T any](src []T, dst []T, equal equalFunc[T]) []T {
	var ret = make([]T, 0, len(src))
	for _, elem := range dst {
		if ContainsFunc[T](src, func(t T) bool {
			return equal(t, elem)
		}) {
			ret = append(ret, elem)
		}
	}

	return deduplicateFunc[T](ret, equal)
}
