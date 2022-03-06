package json_test

import (
	"github.com/rwxrob/fn"
	"github.com/rwxrob/json"
)

func ExampleEscape() {
	set := fn.A[string]{
		`<`, `>`, `&`, `"`, `'`,
		"\t", "\b", "\f", "\n", "\r",
		"\\", "\"", "💢", "д",
	}
	set.Map(json.Escape).Print()
	// Output:
	// <>&\"'\t\b\f\n\r\\\"💢д
}
