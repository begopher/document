package document

func Filter[T any](
	rule Rule[T],
	doc  Document[T],
) Document[T] {
	return filter[T]{rule, doc}
}

type filter[T any] struct {
	rule Rule[T]
	doc  Document[T]
}

func (f filter[T]) Write(value T) error {
	if skip := f.rule.Evaluate(value); skip {
		return nil
	}
	return f.doc.Write(value)
}
