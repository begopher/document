package document

type Rule[T any] interface {
	Evaluate(T) bool
}
