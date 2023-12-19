package slice

import "github.com/gcsdaluo/tekila/internal/Err"

func Add[T any](src []T, element T, index int) ([]T, error) {
	// 先判断越界 再扩展元素
	length := len(src)
	if index < 0 || index >= length {
		return nil, err.ErrIndexOutOfRange(length, index)
	}

	var zeroValue T
	src = append(src, zeroValue)
	for i := len(src) - 1; i > index; i-- {
		if i-1 >= 0 {
			src[i] = src[i-1]
		}
	}
	src[index] = element
	return src, nil
}
