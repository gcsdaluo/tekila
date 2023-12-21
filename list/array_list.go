package list

var (
	_ List[any] = &ArrayList[any]{}
)

type ArrayList[T any] struct {
	elems []T
}
