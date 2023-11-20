package test

import(
	"fmt"
	"testing"
	doc "github.com/begopher/document"
	
)

type Language struct {
	Name string
	Visibility string
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
	err := fmt.Errorf("filterd")
	d := doc.Filter[Language]( 
		_true{},
		err,
		testDoc{t},

	)
	got := d.Write(Language{})
	if err != got {
		t.Errorf("expected error (%v), got (%v)", err, got)
	}
}


func Test_Buffer(t *testing.T) {
	var lang Language
	doc := doc.Buffer(&lang)
	name := "golang"
	visibility := "public"
	err := doc.Write(Language{
		Name: name,
		Visibility: visibility,
	})
	if err != nil {
		t.Errorf("Expected error is (nil) got (%v)", err)
	}
	if lang.Name != name {
		t.Errorf("expected name is (%v) got (%v)", name, lang.Name)
	}
	if lang.Visibility != visibility {
		t.Errorf("expected name is (%v) got (%v)", visibility, lang.Visibility)
	}
}

func Test_Buffers(t *testing.T) {
	langs := make([]Language, 3)
	doc := doc.Buffers(langs)
	// index 0 
	name := "goalng"
	visibility := "public"
	err := doc.Write(Language{
		Name: name,
		Visibility: visibility,
	})
	if err != nil {
		t.Errorf("Expected error is (nil) got (%v)", err)
	}
	lang := langs[0]
	if lang.Name != name {
		t.Errorf("expected name is (%v) got (%v)", name, lang.Name)
	}
	if lang.Visibility != visibility {
		t.Errorf("expected name is (%v) got (%v)", visibility, lang.Visibility)
	}
	// index 1
	name = "any-name"
	visibility = "any-visibility"
	err = doc.Write(Language{
		Name: name,
		Visibility: visibility,
	})
	if err != nil {
		t.Errorf("Expected error is (nil) got (%v)", err)
	}
	lang = langs[1]
	if lang.Name != name {
		t.Errorf("expected name is (%v) got (%v)", name, lang.Name)
	}
	if lang.Visibility != visibility {
		t.Errorf("expected name is (%v) got (%v)", visibility, lang.Visibility)
	}
	// index 2
	lang = langs[2]
	if lang.Name != "" {
		t.Errorf(`expected name is empty string ("") got (%v)`, lang.Name)
	}
	if lang.Visibility != "" {
		t.Errorf(`expected visibility is empty string ("") got (%v)`, lang.Visibility)
	}
}
