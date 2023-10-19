package slice

import "github.com/gcsdaluo/tekila"

// Max 返回最大值。
// 该方法假设你至少会传入一个值。
// 在使用 float32 或者 float64 的时候要小心精度问题
func Max[T tekila.RealNumber](target []T) T {
	res := target[0]
	for i := 1; i < len(target); i++ {
		if target[i] > res {
			res = target[i]
		}
	}
	return res
}

// Min 返回最小值
// 该方法会假设你至少会传入一个值
// 在使用 float32 或者 float64 的时候要小心精度问题
func Min[T tekila.RealNumber](target []T) T {
	res := target[0]
	for i := 1; i < len(target); i++ {
		if target[i] < res {
			res = target[i]
		}
	}
	return res
}

// Sum 求和
// 在使用 float32 或者 float64 的时候要小心精度问题
func Sum[T tekila.Number](target []T) T {
	var res T
	for _, elem := range target {
		res += elem
	}
	return res
}
