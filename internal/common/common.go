package common

func ZeroValueOfT[T any]() T {
	var zero T
	return zero
}
