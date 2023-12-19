package slice

// ReduceMap 执行切片归约操作
// 即重新对字符串进行 连接-长度累加-统计大写字母-整数累乘等操作
func ReduceMap[Src any, Dst any](src []Src, reducer func(Dst, Src) Dst, initialValue Dst) Dst {
	result := initialValue
	for _, elem := range src {
		result = reducer(result, elem)
	}
	return result
}

// FilterMap 执行过滤并且转化
// 如果 m 的第二个返回值是 false，那么我们会忽略第一个返回值
// 即便第二个返回值是 false，后续的元素依旧会被遍历
func FilterMap[Src any, Dst any](src []Src, mapper func(idx int, src Src) (Dst, bool)) []Dst {
	res := make([]Dst, 0, len(src))
	for index, elem := range src {
		dst, ok := mapper(index, elem)
		if ok {
			res = append(res, dst)
		}
	}
	return res
}

// ToMap  功能函数--构造泛型map 将src切片slice转化为一个Map
// map 本身就具备去重功能
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
		if !ContainsFunc[T](data[key+1:], func(dst T) bool {
			return equal(dst, elem)
		}) {
			newData = append(newData, elem)
		}
	}
	return newData
}
