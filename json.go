/*
Package json contains interface specifications for representing any Go
type as JSON where possible. Using the goprintasjson tool allows for quick code generation of scaffolding to make any Go type easily used as JSON.
*/
package json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/rwxrob/fn/each"
	"github.com/rwxrob/to"
)

// AsJSON specifies types that can represent themselves as JSON both in
// a single-line form with no spaces and a long, indented form
// with line returns and consistent 2-space indentation and separation.
type AsJSON interface {
	JSON() ([]byte, error)  // single line, no spaces
	JSONL() ([]byte, error) // 2-space indent and separation
}

// Stringer specifies that rwxrob/io.Stringer interface is fulfilled as
// JSON and adds StringLong for JSONL version. Errors must be logged.
// See AsJSON.
type Stringer interface {
	String() string
	StringLong() string
}

// Printer specifies methods for printing self as JSON and will log any
// error if encountered. Printer provides a consistent representation of
// any structure such that it an easily be read and compared as JSON
// whenever printed and test. Sadly, the default string representations
// for most types in Go are virtually unusable for consistent
// representations of any structure. And while it is true that JSON data
// should be supported in any way that is it presented, some consistent
// output makes for more consistent debugging, documentation, and
// testing.
type Printer interface {
	Print() string
	PrintLong() string
}

// Logger specifies methods for logging self as short and long JSON
// printing separately to the log if any error. The LogLong should print
// a new log entry for each line written.
type Logger interface {
	Log() string
	LogLong() string
}

// Escape only escapes those runes that require it according to the JSON
// specification (unlike the encoding/json standard which defaults to
// escaping many other characters as well unnecessarily).
func Escape(in string) string {
	out := ``
	for _, r := range in {
		switch r {
		case '\t':
			out += `\t`
		case '\b':
			out += `\b`
		case '\f':
			out += `\f`
		case '\n':
			out += `\n`
		case '\r':
			out += `\r`
		case '\\':
			out += `\\`
		case '"':
			out += `\"`
		default:
			out += string(r)
		}
	}
	return out
}

// Marshal mimics json.Marshal from the encoding/json package without
// the broken, unnecessary HTML escapes and extraneous newline that the
// json.Encoder adds. Call this from your own MarshalJSON methods to get
// JSON rendering that is more readable and compliant with the JSON
// specification (unless you are using the extremely rare case of
// dumping that into HTML, for some reason).
func Marshal(v any) ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	err := enc.Encode(v)
	return []byte(strings.TrimSpace(buf.String())), err
}

// Unmarshal mimics json.Unmarshal from the encoding/json package.
func Unmarshal(buf []byte, v any) error {
	return json.Unmarshal(buf, v)
}

// ------------------------------- Array ------------------------------

// Array is a slice of strings that knows how to marshal as JSON.
type Array []string

// JSONL implements rwxrob/json.AsJSON.
func (s Array) JSON() ([]byte, error) { return json.Marshal(s) }

// JSONL implements rwxrob/json.AsJSON.
func (s Array) JSONL() ([]byte, error) {
	return json.MarshalIndent(s, "  ", "  ")
}

// String implements rwxrob/json.Stringer and fmt.Stringer.
func (s Array) String() string {
	byt, err := s.JSON()
	if err != nil {
		log.Print(err)
	}
	return string(byt)
}

// StringLong implements rwxrob/json.Stringer.
func (s Array) StringLong() string {
	byt, err := s.JSONL()
	if err != nil {
		log.Print(err)
	}
	return string(byt)
}

// String implements rwxrob/json.Printer.
func (s Array) Print() { fmt.Println(s.String()) }

// PrintLong implements rwxrob/json.Printer.
func (s Array) PrintLong() { fmt.Println(s.StringLong()) }

// Log implements rwxrob/json.Logger.
func (s Array) Log() { log.Print(s.String()) }

// LogLong implements rwxrob/json.Logger.
func (s Array) LogLong() { each.Log(to.Lines(s.StringLong())) }

// ------------------------------ Object ------------------------------

// Object represents any JSON-able struct
type Object struct{ This any }

// MarshalJSON implements rwxrob/json.AsJSON
func (s Object) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.This)
}

// UnmarshalJSON implements rwxrob/json.AsJSON
func (s *Object) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &s.This)
}

// JSONL implements rwxrob/json.AsJSON.
func (s Object) JSON() ([]byte, error) { return json.Marshal(s) }

// JSONL implements rwxrob/json.AsJSON.
func (s Object) JSONL() ([]byte, error) {
	return json.MarshalIndent(s, "  ", "  ")
}

// String implements rwxrob/json.Stringer and fmt.Stringer.
func (s Object) String() string {
	byt, err := s.JSON()
	if err != nil {
		log.Print(err)
	}
	return string(byt)
}

// StringLong implements rwxrob/json.Stringer.
func (s Object) StringLong() string {
	byt, err := s.JSONL()
	if err != nil {
		log.Print(err)
	}
	return string(byt)
}

// String implements rwxrob/json.Printer.
func (s Object) Print() { fmt.Println(s.String()) }

// PrintLong implements rwxrob/json.Printer.
func (s Object) PrintLong() { fmt.Println(s.StringLong()) }

// Log implements rwxrob/json.Logger.
func (s Object) Log() { log.Print(s.String()) }

// LogLong implements rwxrob/json.Logger.
func (s Object) LogLong() { each.Log(to.Lines(s.StringLong())) }
