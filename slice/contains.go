package slice

// Contains 判断src中是否包含dst
func Contains[T comparable](src []T, dst T) bool {
	return ContainsFunc[T](src, func(elem T) bool {
		return elem == dst
	})
}

// ContainsFunc 判断src中是否包含dst
// 通过传入自定义比较函数进行比较
func ContainsFunc[T any](src []T, equal func(dst T) bool) bool {
	for _, elem := range src {
		if equal(elem) {
			return true
		}
	}
	return false
}

// ContainsOne 判断src中是否存在dst的任意一元素
func ContainsOne[T comparable](src, dst []T) bool {
	srcMapping := sToMap[T](src)
	for _, elem := range dst {
		if _, exist := srcMapping[elem]; exist {
			return true
		}
	}
	return false
}

// ContainsOneFunc 判断src中是否包含dst的任一元素
// 通过传入自定义比较函数进行比较
func ContainsOneFunc[T any](src, dst []T, equal equalFunc[T]) bool {
	for _, elemDst := range dst {
		for _, elemSrc := range src {
			if equal(elemSrc, elemDst) {
				return true
			}
		}
	}
	return false
}

// ContainsAll 判断src里面是否存在dst的所有元素
func ContainsAll[T comparable](src, dst []T) bool {
	srcMapping := sToMap[T](src)
	for _, elem := range dst {
		if _, exist := srcMapping[elem]; !exist {
			return false
		}
	}
	return true
}

// ContainsAllFunc 判断src里面是否存在dst的所有元素
// 通过传入自定义比较函数进行比较
func ContainsAllFunc[T any](src, dst []T, equal equalFunc[T]) bool {
	for _, elemDst := range dst {
		// 使用自定义比较函数检查 elemDst 是否存在于 src 中
		if !ContainsFunc[T](src, func(elem T) bool {
			return equal(elem, elemDst)
		}) {
			return false
		}
	}
	return true
}
