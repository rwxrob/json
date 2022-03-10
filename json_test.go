package json_test

import (
	"fmt"
	"log"
	"os"

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

func ExampleArray_JSON() {
	list := json.Array([]string{"foo", "bar"})
	byt, _ := list.JSON()
	fmt.Println(string(byt))
	// Output:
	// ["foo","bar"]
}

func ExampleArray_JSONL() {
	list := json.Array([]string{"foo", "bar"})
	byt, _ := list.JSONL()
	fmt.Println(string(byt))
	// Output:
	// [
	//     "foo",
	//     "bar"
	//   ]
}

func ExampleArray_String() {
	list := json.Array([]string{"foo", "bar"})
	fmt.Println(list.String())
	// Output:
	// ["foo","bar"]
}

func ExampleArray_StringLong() {
	list := json.Array([]string{"foo", "bar"})
	fmt.Println(list.StringLong())
	// Output:
	// [
	//     "foo",
	//     "bar"
	//   ]
}

func ExampleArray_Print() {
	list := json.Array([]string{"foo", "bar"})
	list.Print()
	// Output:
	// ["foo","bar"]
}

func ExampleArray_PrintLong() {
	list := json.Array([]string{"foo", "bar"})
	list.PrintLong()

	//list.LogLong() // also check this to be sure one line per log line

	// Output:
	// [
	//     "foo",
	//     "bar"
	//   ]
}

func ExampleArray_Log() {

	// adjust log output for testing
	log.SetOutput(os.Stdout)
	log.SetFlags(0)
	defer log.SetOutput(os.Stderr)
	defer log.SetFlags(log.Flags())

	list := json.Array([]string{"foo", "bar"})
	list.Log()

	// Output:
	// ["foo","bar"]
}

func ExampleArray_LogLong() {

	// adjust log output for testing
	log.SetOutput(os.Stdout)
	log.SetFlags(0)
	defer log.SetOutput(os.Stderr)
	defer log.SetFlags(log.Flags())

	list := json.Array([]string{"foo", "bar"})
	list.LogLong()

	// Output:
	// [
	//     "foo",
	//     "bar"
	//   ]
}

func ExampleObject_Print() {

	type FooBar struct {
		Foo string
		Bar string
	}

	json.Object[FooBar]{FooBar{"FOO", "BAR"}}.Print()

	// Output:
	// {"Foo":"FOO","Bar":"BAR"}
}
