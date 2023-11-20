package document

func Buffer[T any](buff *T) Document[T] {
	if buff == nil {
		panic("document.Buffer: buff cannot be nil")
	}
	return buffer[T]{buff}
}

type buffer[T any] struct {
	buff *T
}

func (b buffer[T]) Write(value T) error {
	*b.buff = value
	return nil
}
