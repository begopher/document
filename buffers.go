package document

func Buffers[T any](buff []T) Document[T] {
	if buff == nil {
		panic("document.Buffers: buff cannot be nil")
	}
	if len(buff) == 0 {
		panic("document.Buffers: buff's length cannot be zero")
	}
	return &buffers[T]{buff: buff}
}

type buffers[T any] struct {
	index int
	buff []T
}

func (b *buffers[T]) Write(value T) error {
	if b.index >= len(b.buff) {
		panic("document.buffers.Write: buffer is full")
	}
	b.buff[b.index] = value
	b.index++
	return nil
}
