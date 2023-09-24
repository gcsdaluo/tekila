package slice

// ToMap  功能函数--构造泛型map 将src切片slice转化为一个Map
func sToMap[T comparable](src []T) map[T]struct{} {
	var dataMap = make(map[T]struct{}, len(src))
	for _, elem := range src {
		dataMap[elem] = struct{}{} // 使用空结构体可以减少内存消耗
	}

	return dataMap
}

// deduplicate 功能函数--去除重复元素 返回去重切片
func deduplicate[T comparable](data []T) []T {
	dataMap := sToMap[T](data) // 去重
	var newData = make([]T, 0, len(dataMap))
	for key := range dataMap {
		newData = append(newData, key)
	}
	return newData
}

// deduplicateFunc 功能函数--去除重复元素 返回去重切片
// 通过自定义比较函数
func deduplicateFunc[T any](data []T, equal equalFunc[T]) []T {
	var newData = make([]T, 0, len(data))
	for key, elem := range data {
		if !ContainsFunc[T](data[key+1:], elem, equal) {
			newData = append(newData, elem)
		}
	}
	return newData
}
