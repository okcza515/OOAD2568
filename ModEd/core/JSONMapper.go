package core

type JSONMapper[T any] struct {
	Path string
}

func (mapper *JSONMapper[T]) Deserialize() []*T {
	result := []*T{}
	return result
}
