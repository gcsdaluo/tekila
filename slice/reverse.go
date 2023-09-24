package slice

// Reverse 创建新切片进行切片翻转
func Reverse[T comparable](src []T) []T {
	var target = make([]T, 0, len(src))
	for i := len(src) - 1; i >= 0; i-- {
		target = append(target, src[i])
	}
	return target
}

// ReverseOrigin 在原切片进行翻转
func ReverseOrigin[T comparable](src []T) {
	for i, j := 0, len(src)-1; i < j; i, j = i+1, j-1 {
		src[i], src[j] = src[j], src[i]
	}
}
