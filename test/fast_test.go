package test

import(
	"testing"
	doc "github.com/begopher/document"
	
)

type Language struct {
	Name string
	Visibility string
	Rank int
}


type testDoc struct {
	t *testing.T
}

func (tc testDoc) Write(lang Language) error {
	tc.t.Errorf("Not filtered")
	return nil
}

type _true struct {}
func (_true) Evaluate(Language) bool {
	return true
}

type _false struct {}
func (_false) Evaluate(Language) bool {
	return false
}

func Test_Filter(t *testing.T) {
	d := doc.Filter[Language]( 
		_true{},
		testDoc{t},

	)
	d.Write(Language{})	
}
