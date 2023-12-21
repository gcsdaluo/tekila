package list

// List 接口 仅定义方法的行为和表现
type List[T any] interface {
	Get(index int) (T, error)
	Set(index int, target T) error
	Add(index int, target T) error
	Delete(index int) (T, error)
	Append(targets ...T) error
	Len() int
	Cap() int
	Range(fn func(index int, target T) error) error
	ToSlice() []T
}
