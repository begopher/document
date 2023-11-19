package document

type Document[T any] interface {
	Write(T) error
}


