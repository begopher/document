package document

func Filter[T any](
	rule Rule[T],
	err error,
	doc  Document[T],
) Document[T] {
	if rule == nil {
		panic("document.Filter: rule cannot be nil")
	}
	if err == nil {
		panic("document.Filter: err cannot be nil")
	}
	if doc == nil {
		panic("document.Filter: doc cannot be nil")
	}
	return filter[T]{rule, err, doc}
}

type filter[T any] struct {
	rule Rule[T]
	err error
	doc  Document[T]
}

func (f filter[T]) Write(value T) error {
	if skip := f.rule.Evaluate(value); skip {
		return f.err
	}
	return f.doc.Write(value)
}
